package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Items")
	vars := mux.Vars(r)

	items := models.GetItems(vars["id"])
	items.Encode(w)
}

func (handler *ItemHandler) GetLists(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Lists")
	vars := mux.Vars(r)

	items := models.GetLists(vars["id"])
	items.Encode(w)
}
