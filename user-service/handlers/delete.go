package handlers

import (
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/user-service/authentication"
	"github.com/v24Zer0/ToDo/user-service/database"
	"github.com/v24Zer0/ToDo/user-service/models"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete request - User")

	var user models.UserNoPassword
	err := user.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	c := make(chan authentication.UserValidation)
	token := r.Header.Get("Token")

	go authentication.ValidateToken(c, token)

	userValid := <-c
	if userValid.Error != nil {
		http.Error(w, "not authenticated", http.StatusUnauthorized)
		return
	}

	if userValid.UserID != user.ID {
		http.Error(w, "cannot delete another user", http.StatusForbidden)
		return
	}

	err = database.DeleteUser(h.db, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
