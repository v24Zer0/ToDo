package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/auth-service/database"
)

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token, err := database.RetrieveToken(h.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token.Encode(w)
}
