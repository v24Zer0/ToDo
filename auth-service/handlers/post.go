package handlers

import (
	"net/http"

	"github.com/v24Zer0/ToDo/auth-service/database"
	"github.com/v24Zer0/ToDo/auth-service/models"
)

func (h *Handler) CreateToken(w http.ResponseWriter, r *http.Request) {
	var user models.UserToken
	err := user.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.UserID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	err = database.CreateToken(h.db, user.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var token models.UserToken
	err := token.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if token.Token == "" {
		http.Error(w, "missing token", http.StatusBadRequest)
		return
	}

	t, err := database.RetrieveByToken(h.db, token.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t.Encode(w)
}
