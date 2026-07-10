package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DangDDT/tada-learn-english/backend/internal/middleware"
	"github.com/DangDDT/tada-learn-english/backend/internal/service"
	"github.com/go-chi/chi/v5"
)

type WordHandler struct {
	svc *service.WordService
}

func NewWordHandler(svc *service.WordService) *WordHandler {
	return &WordHandler{svc: svc}
}

func (h *WordHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var input service.CreateWordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid request body")
		return
	}
	if input.Word == "" || input.Meaning == "" {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Word and meaning required")
		return
	}

	word, err := h.svc.Create(r.Context(), userID, input)
	if err != nil {
		if err == service.ErrWordDuplicate {
			writeError(w, http.StatusConflict, "DUPLICATE_WORD", err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create word")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    word,
	})
}

func (h *WordHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page < 1 {
		page = 1
	}
	perPage, _ := strconv.Atoi(q.Get("per_page"))
	if perPage < 1 || perPage > 100 {
		perPage = 20
	}

	params := service.WordListParams{
		UserID:    userID,
		Query:     q.Get("q"),
		CEFRLevel: q.Get("cefr_level"),
		SRSBand:   q.Get("srs_band"),
		Tag:       q.Get("tag"),
		SortBy:    q.Get("sort_by"),
		SortDir:   q.Get("sort_dir"),
		Page:      page,
		PerPage:   perPage,
	}

	result, err := h.svc.List(r.Context(), params)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to list words")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    result.Words,
		"meta":    result.Meta,
	})
}

func (h *WordHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	wordID := chi.URLParam(r, "id")

	word, err := h.svc.Get(r.Context(), userID, wordID)
	if err != nil {
		if err == service.ErrWordNotFound {
			writeError(w, http.StatusNotFound, "NOT_FOUND", "Word not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get word")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    word,
	})
}

func (h *WordHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	wordID := chi.URLParam(r, "id")

	var input service.UpdateWordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Invalid request body")
		return
	}

	word, err := h.svc.Update(r.Context(), userID, wordID, input)
	if err != nil {
		if err == service.ErrWordNotFound {
			writeError(w, http.StatusNotFound, "NOT_FOUND", "Word not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to update word")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    word,
	})
}

func (h *WordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	wordID := chi.URLParam(r, "id")

	if err := h.svc.Delete(r.Context(), userID, wordID); err != nil {
		if err == service.ErrWordNotFound {
			writeError(w, http.StatusNotFound, "NOT_FOUND", "Word not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to delete word")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *WordHandler) Import(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	file, _, err := r.FormFile("csv")
	if err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "CSV file required")
		return
	}
	defer file.Close()

	result, err := h.svc.Import(r.Context(), userID, file)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to import words")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    result,
	})
}
