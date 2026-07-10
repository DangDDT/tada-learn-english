package service

import (
	"context"
	"encoding/csv"
	"errors"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/DangDDT/tada-learn-english/backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrWordNotFound = errors.New("word not found")
	ErrWordDuplicate = errors.New("word exists")
)

type WordService struct {
	pool *pgxpool.Pool
}

func NewWordService(pool *pgxpool.Pool) *WordService {
	return &WordService{pool: pool}
}

type CreateWordInput struct {
	Word             string   `json:"word"`
	Pronunciation    string   `json:"pronunciation,omitempty"`
	IPA              string   `json:"ipa,omitempty"`
	Meaning          string   `json:"meaning"`
	PartOfSpeech     string   `json:"part_of_speech,omitempty"`
	ExampleSentences []string `json:"example_sentences,omitempty"`
	CEFRLevel        string   `json:"cefr_level,omitempty"`
	Tags             []string `json:"tags,omitempty"`
}

type UpdateWordInput struct {
	Word             *string   `json:"word,omitempty"`
	Pronunciation    *string   `json:"pronunciation,omitempty"`
	IPA              *string   `json:"ipa,omitempty"`
	Meaning          *string   `json:"meaning,omitempty"`
	PartOfSpeech     *string   `json:"part_of_speech,omitempty"`
	ExampleSentences *[]string `json:"example_sentences,omitempty"`
	CEFRLevel        *string   `json:"cefr_level,omitempty"`
	Tags             *[]string `json:"tags,omitempty"`
}

type WordListParams struct {
	UserID    string
	Query     string
	CEFRLevel string
	SRSBand   string
	Tag       string
	SortBy    string
	SortDir   string
	Page      int
	PerPage   int
}

type WordListResponse struct {
	Words []model.WordWithSRS `json:"data"`
	Meta  struct {
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
		Total   int `json:"total"`
	} `json:"meta"`
}

type ImportResult struct {
	Imported int              `json:"imported"`
	Skipped  int              `json:"skipped"`
	Errors   []ImportErrorRow `json:"errors"`
}

type ImportErrorRow struct {
	Row   int    `json:"row"`
	Word  string `json:"word"`
	Error string `json:"error"`
}

func (s *WordService) Create(ctx context.Context, userID string, input CreateWordInput) (*model.WordWithSRS, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	word := &model.WordWithSRS{}
	err = tx.QueryRow(ctx, `
		WITH new_word AS (
			INSERT INTO words (user_id, word, pronunciation, ipa, meaning, part_of_speech,
			                   example_sentences, cefr_level, tags)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (user_id, word, deleted_at) DO NOTHING
			RETURNING id, user_id, word, pronunciation, ipa, meaning, part_of_speech,
			          example_sentences, cefr_level, tags, created_at, updated_at
		), new_srs AS (
			INSERT INTO srs_states (word_id, user_id)
			SELECT id, user_id FROM new_word
			RETURNING band, times_reviewed, last_reviewed_at, next_review_at
		)
		SELECT w.id, w.word, w.pronunciation, w.ipa, w.meaning, w.part_of_speech,
		       w.example_sentences, w.cefr_level, w.tags, w.created_at, w.updated_at,
		       s.band, s.times_reviewed, s.last_reviewed_at, s.next_review_at
		FROM new_word w, new_srs s`,
		userID, input.Word, input.Pronunciation, input.IPA, input.Meaning,
		input.PartOfSpeech, input.ExampleSentences, input.CEFRLevel, input.Tags,
	).Scan(
		&word.ID, &word.Word, &word.Pronunciation, &word.IPA, &word.Meaning,
		&word.PartOfSpeech, &word.ExampleSentences, &word.CEFRLevel, &word.Tags,
		&word.CreatedAt, &word.UpdatedAt, &word.SRSBand, &word.TimesReviewed,
		&word.SRSLastReview, &word.SRSNextReview,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWordDuplicate
		}
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return word, nil
}

func (s *WordService) List(ctx context.Context, params WordListParams) (*WordListResponse, error) {
	where := []string{"w.user_id = $1", "w.deleted_at IS NULL"}
	args := []interface{}{params.UserID}
	argIdx := 2

	if params.Query != "" {
		where = append(where, "(w.word ILIKE $"+itos(argIdx)+" OR w.meaning ILIKE $"+itos(argIdx)+")")
		args = append(args, "%"+params.Query+"%")
		argIdx++
	}
	if params.CEFRLevel != "" {
		where = append(where, "w.cefr_level = $"+itos(argIdx))
		args = append(args, params.CEFRLevel)
		argIdx++
	}
	if params.SRSBand != "" {
		where = append(where, "s.band = $"+itos(argIdx))
		args = append(args, params.SRSBand)
		argIdx++
	}
	if params.Tag != "" {
		where = append(where, "$"+itos(argIdx)+" = ANY(w.tags)")
		args = append(args, params.Tag)
		argIdx++
	}

	if params.SortBy == "" {
		params.SortBy = "created_at"
	}
	if params.SortDir == "" {
		params.SortDir = "DESC"
	}
	validSorts := map[string]bool{"created_at": true, "updated_at": true, "word": true, "cefr_level": true}
	if !validSorts[params.SortBy] {
		params.SortBy = "created_at"
	}
	if params.SortDir != "ASC" && params.SortDir != "DESC" {
		params.SortDir = "DESC"
	}

	countQuery := "SELECT COUNT(*) FROM words w LEFT JOIN srs_states s ON s.word_id = w.id AND s.user_id = w.user_id WHERE " + strings.Join(where, " AND ")
	total := 0
	err := s.pool.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	offset := (params.Page - 1) * params.PerPage
	limitArg := "$" + itos(argIdx)
	args = append(args, params.PerPage)
	argIdx++
	offsetArg := "$" + itos(argIdx)
	args = append(args, offset)

	query := `SELECT w.id, w.word, w.pronunciation, w.ipa, w.meaning, w.part_of_speech,
	       w.example_sentences, w.cefr_level, w.tags, w.created_at, w.updated_at,
	       s.band, s.times_reviewed, s.last_reviewed_at, s.next_review_at
	FROM words w
	LEFT JOIN srs_states s ON s.word_id = w.id AND s.user_id = w.user_id
	WHERE ` + strings.Join(where, " AND ") + `
	ORDER BY w.` + params.SortBy + " " + params.SortDir + `
	LIMIT ` + limitArg + ` OFFSET ` + offsetArg

	rows, err := s.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	words := make([]model.WordWithSRS, 0)
	for rows.Next() {
		var w model.WordWithSRS
		err := rows.Scan(
			&w.ID, &w.Word, &w.Pronunciation, &w.IPA, &w.Meaning,
			&w.PartOfSpeech, &w.ExampleSentences, &w.CEFRLevel, &w.Tags,
			&w.CreatedAt, &w.UpdatedAt, &w.SRSBand, &w.TimesReviewed,
			&w.SRSLastReview, &w.SRSNextReview,
		)
		if err != nil {
			return nil, err
		}
		words = append(words, w)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	resp := &WordListResponse{
		Words: words,
		Meta: struct {
			Page    int `json:"page"`
			PerPage int `json:"per_page"`
			Total   int `json:"total"`
		}{
			Page:    params.Page,
			PerPage: params.PerPage,
			Total:   total,
		},
	}
	return resp, nil
}

func (s *WordService) Get(ctx context.Context, userID, wordID string) (*model.WordWithSRS, error) {
	word := &model.WordWithSRS{}
	err := s.pool.QueryRow(ctx, `
		SELECT w.id, w.word, w.pronunciation, w.ipa, w.meaning, w.part_of_speech,
		       w.example_sentences, w.cefr_level, w.tags, w.created_at, w.updated_at,
		       s.band, s.times_reviewed, s.last_reviewed_at, s.next_review_at
		FROM words w
		LEFT JOIN srs_states s ON s.word_id = w.id AND s.user_id = w.user_id
		WHERE w.id = $1 AND w.user_id = $2 AND w.deleted_at IS NULL`,
		wordID, userID,
	).Scan(
		&word.ID, &word.Word, &word.Pronunciation, &word.IPA, &word.Meaning,
		&word.PartOfSpeech, &word.ExampleSentences, &word.CEFRLevel, &word.Tags,
		&word.CreatedAt, &word.UpdatedAt, &word.SRSBand, &word.TimesReviewed,
		&word.SRSLastReview, &word.SRSNextReview,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWordNotFound
		}
		return nil, err
	}
	return word, nil
}

func (s *WordService) Update(ctx context.Context, userID, wordID string, input UpdateWordInput) (*model.WordWithSRS, error) {
	var ownerID string
	err := s.pool.QueryRow(ctx, "SELECT user_id FROM words WHERE id=$1 AND deleted_at IS NULL", wordID).Scan(&ownerID)
	if err != nil {
		return nil, ErrWordNotFound
	}
	if ownerID != userID {
		return nil, ErrWordNotFound
	}

	if input.Word != nil {
		var exists bool
		err = s.pool.QueryRow(ctx,
			"SELECT EXISTS(SELECT 1 FROM words WHERE user_id=$1 AND word=$2 AND id != $3 AND deleted_at IS NULL)",
			userID, *input.Word, wordID,
		).Scan(&exists)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ErrWordDuplicate
		}
	}

	setClauses := []string{}
	args := []interface{}{}
	argIdx := 1

	if input.Word != nil {
		setClauses = append(setClauses, "word=$"+itos(argIdx))
		args = append(args, *input.Word)
		argIdx++
	}
	if input.Pronunciation != nil {
		setClauses = append(setClauses, "pronunciation=$"+itos(argIdx))
		args = append(args, *input.Pronunciation)
		argIdx++
	}
	if input.IPA != nil {
		setClauses = append(setClauses, "ipa=$"+itos(argIdx))
		args = append(args, *input.IPA)
		argIdx++
	}
	if input.Meaning != nil {
		setClauses = append(setClauses, "meaning=$"+itos(argIdx))
		args = append(args, *input.Meaning)
		argIdx++
	}
	if input.PartOfSpeech != nil {
		setClauses = append(setClauses, "part_of_speech=$"+itos(argIdx))
		args = append(args, *input.PartOfSpeech)
		argIdx++
	}
	if input.ExampleSentences != nil {
		setClauses = append(setClauses, "example_sentences=$"+itos(argIdx))
		args = append(args, *input.ExampleSentences)
		argIdx++
	}
	if input.CEFRLevel != nil {
		setClauses = append(setClauses, "cefr_level=$"+itos(argIdx))
		args = append(args, *input.CEFRLevel)
		argIdx++
	}
	if input.Tags != nil {
		setClauses = append(setClauses, "tags=$"+itos(argIdx))
		args = append(args, *input.Tags)
		argIdx++
	}

	if len(setClauses) == 0 {
		return s.Get(ctx, userID, wordID)
	}

	setClauses = append(setClauses, "updated_at=NOW()")
	query := "UPDATE words SET " + strings.Join(setClauses, ", ") + " WHERE id=$"+itos(argIdx)
	args = append(args, wordID)

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return s.Get(ctx, userID, wordID)
}

func (s *WordService) Delete(ctx context.Context, userID, wordID string) error {
	tag, err := s.pool.Exec(ctx,
		"UPDATE words SET deleted_at = NOW() WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL",
		wordID, userID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrWordNotFound
	}
	return nil
}

func (s *WordService) Import(ctx context.Context, userID string, file multipart.File) (*ImportResult, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	result := &ImportResult{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 2 {
			result.Errors = append(result.Errors, ImportErrorRow{Row: i + 1, Error: "invalid format: need at least word and meaning"})
			result.Skipped++
			continue
		}

		word := record[0]
		meaning := record[1]
		var partOfSpeech, cefrLevel string
		var exampleSentences []string
		if len(record) > 2 {
			partOfSpeech = record[2]
		}
		if len(record) > 3 {
			cefrLevel = record[3]
		}
		if len(record) > 4 {
			exampleSentences = []string{record[4]}
		}

		var wordID string
		err = tx.QueryRow(ctx, `
			INSERT INTO words (user_id, word, pronunciation, ipa, meaning, part_of_speech,
			                   example_sentences, cefr_level, tags)
			VALUES ($1, $2, '', '', $3, $4, $5, $6, '{}')
			ON CONFLICT (user_id, word, deleted_at) DO NOTHING
			RETURNING id`,
			userID, word, meaning, partOfSpeech, exampleSentences, cefrLevel,
		).Scan(&wordID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				result.Skipped++
				continue
			}
			return nil, err
		}

		result.Imported++
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

func itos(i int) string {
	return strconv.Itoa(i)
}