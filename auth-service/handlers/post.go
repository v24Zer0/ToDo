package handlers

import "net/http"

func (h *Handler) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
