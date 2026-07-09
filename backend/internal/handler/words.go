package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DangDDT/tada-learn-english/backend/internal/service"
)

type WordHandler struct {
	svc *service.WordService
}

func NewWordHandler(svc *service.WordService) *WordHandler {
	return &WordHandler{svc: svc}
}

func (h *WordHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input service.CreateWordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("user_id").(string)
	word, err := h.svc.Create(r.Context(), userID, input)
	if err != nil {
		if err == service.ErrWordDuplicate {
			http.Error(w, `{"error":"word already exists"}`, http.StatusConflict)
			return
		}
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(word)
}

func (h *WordHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	params := service.WordListParams{
		UserID:    userID,
		Query:     r.URL.Query().Get("q"),
		CEFRLevel: r.URL.Query().Get("cefr_level"),
		SRSBand:   r.URL.Query().Get("srs_band"),
		Tag:       r.URL.Query().Get("tag"),
		SortBy:    r.URL.Query().Get("sort_by"),
		SortDir:   r.URL.Query().Get("sort_dir"),
		Page:      1,
		PerPage:   20,
	}

	resp, err := h.svc.List(r.Context(), params)
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *WordHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	wordID := r.PathValue("id")

	word, err := h.svc.Get(r.Context(), userID, wordID)
	if err != nil {
		if err == service.ErrWordNotFound {
			http.Error(w, `{"error":"word not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

func (h *WordHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	wordID := r.PathValue("id")

	var input service.UpdateWordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	word, err := h.svc.Update(r.Context(), userID, wordID, input)
	if err != nil {
		if err == service.ErrWordNotFound {
			http.Error(w, `{"error":"word not found"}`, http.StatusNotFound)
			return
		}
		if err == service.ErrWordDuplicate {
			http.Error(w, `{"error":"word already exists"}`, http.StatusConflict)
			return
		}
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

func (h *WordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	wordID := r.PathValue("id")

	if err := h.svc.Delete(r.Context(), userID, wordID); err != nil {
		if err == service.ErrWordNotFound {
			http.Error(w, `{"error":"word not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *WordHandler) Import(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, `{"error":"file required"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := h.svc.Import(r.Context(), userID, file)
	if err != nil {
		http.Error(w, `{"error":"import failed"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
