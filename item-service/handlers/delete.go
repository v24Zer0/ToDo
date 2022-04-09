package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/item-service/database"
)

func (handler *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete request - Item")

	vars := mux.Vars(r)

	err := database.DeleteItem(handler.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) DeleteList(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete request - List")

	vars := mux.Vars(r)

	err := database.DeleteList(handler.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
