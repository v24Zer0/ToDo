package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Get request - Items")

	items := models.GetItems(vars["id"])
	items.Encode(w)
}
