package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/item-service/database"
	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Put request - Item")

	var item models.Item
	err := item.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUpdateItem(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateItem(handler.db, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) UpdateList(w http.ResponseWriter, r *http.Request) {
	log.Println("Put request - List")

	var list models.List
	err := list.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUpdateList(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateList(handler.db, &list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateUpdateItem(item *models.Item) error {
	if item.ID == "" {
		return errors.New("missing item id")
	}

	if item.Task == "" {
		return errors.New("missing item task")
	}
	return nil
}

func validateUpdateList(list *models.List) error {
	if list.ID == "" {
		return errors.New("missing list id")
	}

	if list.Name == "" {
		return errors.New("missing list name")
	}
	return nil
}
