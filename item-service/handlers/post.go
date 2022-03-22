package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - Item")

	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println("Could not decode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateItemRequest(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(item.Task)

	w.WriteHeader(http.StatusOK)
}

func (handler *ItemHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - List")

	var list models.List
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		log.Println("Could not decode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validateListRequest(&list)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(list.Name)

	w.WriteHeader(http.StatusOK)
}

func validateItemRequest(item *models.Item) error {
	if item.Task == "" {
		return errors.New("missing task")
	}
	if item.ListID == "" {
		return errors.New("missing list_id")
	}
	return nil
}

func validateListRequest(list *models.List) error {
	if list.Name == "" {
		return errors.New("missing name")
	}

	if list.UserID == "" {
		return errors.New("missing user_id")
	}
	return nil
}
