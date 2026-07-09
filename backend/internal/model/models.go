package model

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Word struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	Word             string    `json:"word"`
	Pronunciation    string    `json:"pronunciation,omitempty"`
	IPA              string    `json:"ipa,omitempty"`
	Meaning          string    `json:"meaning"`
	PartOfSpeech     string    `json:"part_of_speech,omitempty"`
	ExampleSentences []string  `json:"example_sentences,omitempty"`
	CEFRLevel        string    `json:"cefr_level,omitempty"`
	Tags             []string  `json:"tags,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type WordWithSRS struct {
	ID               string    `json:"id"`
	Word             string    `json:"word"`
	Pronunciation    string    `json:"pronunciation,omitempty"`
	IPA              string    `json:"ipa,omitempty"`
	Meaning          string    `json:"meaning"`
	PartOfSpeech     string    `json:"part_of_speech,omitempty"`
	ExampleSentences []string  `json:"example_sentences,omitempty"`
	CEFRLevel        string    `json:"cefr_level,omitempty"`
	Tags             []string  `json:"tags,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	SRSBand          string    `json:"srs_band"`
	TimesReviewed    int       `json:"times_reviewed"`
	SRSLastReview    *time.Time `json:"srs_last_review,omitempty"`
	SRSNextReview    *time.Time `json:"srs_next_review,omitempty"`
}

type SRSState struct {
	ID            string    `json:"id"`
	WordID        string    `json:"word_id"`
	UserID        string    `json:"user_id"`
	Band          string    `json:"band"`
	TimesReviewed int       `json:"times_reviewed"`
	LastReviewed  *time.Time `json:"last_reviewed_at,omitempty"`
	NextReview    *time.Time `json:"next_review_at,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
