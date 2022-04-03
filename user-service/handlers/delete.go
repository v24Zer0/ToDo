package handlers

import (
	"log"
	"net/http"

	"github.com/v24Zer0/ToDO/user-service/database"
	"github.com/v24Zer0/ToDO/user-service/models"
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

	err = database.DeleteUser(h.db, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
