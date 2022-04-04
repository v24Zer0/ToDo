package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"github.com/v24Zer0/ToDo/item-service/database"
)

func (handler *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Items")
	vars := mux.Vars(r)

	items, err := database.RetrieveItems(handler.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addHeaders(w)
	items.Encode(w)
}

func (handler *Handler) GetLists(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Lists")
	vars := mux.Vars(r)

	items, err := database.RetrieveLists(handler.db, vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	addHeaders(w)
	items.Encode(w)
}

func (handler *Handler) GetID(w http.ResponseWriter, r *http.Request) {
	id := ksuid.New()
	w.Write([]byte(id.String()))
}
