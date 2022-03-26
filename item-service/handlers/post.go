package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/v24Zer0/ToDo/item-service/database"
	"github.com/v24Zer0/ToDo/item-service/models"
)

func (handler *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - Item")

	var item models.Item
	err := item.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validatePostItem(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateItem(handler.db, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *ItemHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	log.Println("Post request - List")

	var list models.List
	err := list.Decode(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validatePostList(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.CreateList(handler.db, &list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validatePostItem(item *models.Item) error {
	if item.Task == "" {
		return errors.New("missing item task")
	}

	if item.ListID == "" {
		return errors.New("missing list_id")
	}
	return nil
}

func validatePostList(list *models.List) error {
	if list.Name == "" {
		return errors.New("missing list name")
	}

	if list.UserID == "" {
		return errors.New("missing user_id")
	}
	return nil
}
