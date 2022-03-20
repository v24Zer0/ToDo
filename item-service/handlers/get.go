package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

func (handler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Items")
	vars := mux.Vars(r)

	items := handler.db.RetrieveItems(vars["id"])
	items.Encode(w)
}

func (handler *ItemHandler) GetLists(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Lists")
	vars := mux.Vars(r)

	items := handler.db.RetrieveLists(vars["id"])
	items.Encode(w)
}

func (handler *ItemHandler) GetID(w http.ResponseWriter, r *http.Request) {
	id := ksuid.New()
	w.Write([]byte(id.String()))
}
