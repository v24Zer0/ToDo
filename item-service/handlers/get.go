package handlers

import (
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	log.Println("Get request - Items")
	items := models.GetItems()
	items.Encode(w)
}
