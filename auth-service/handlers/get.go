package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/auth-service/database"
)

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Token")

	vars := mux.Vars(r)

	token, err := database.RetrieveToken(h.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token.Encode(w)
}
