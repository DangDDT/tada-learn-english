package handler

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	pool interface {
		Ping() error
	}
}

func NewHealthHandler(pool interface{ Ping() error }) *HealthHandler {
	return &HealthHandler{pool: pool}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	dbStatus := "connected"
	if err := h.pool.Ping(); err != nil {
		dbStatus = "disconnected"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"version": "1.0.0",
		"db":      dbStatus,
	})
}
