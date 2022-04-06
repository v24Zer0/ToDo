package handlers

import (
	"errors"
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

	err = validateTokenRequest(&token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := database.RetrieveToken(h.db, token.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.Token != token.Token {
		http.Error(w, "token does not match", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateTokenRequest(token *models.UserToken) error {
	if token.UserID == "" {
		return errors.New("missing user_id")
	}

	if token.Token == "" {
		return errors.New("missing token")
	}
	return nil
}
