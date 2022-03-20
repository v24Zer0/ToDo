package handlers

import "github.com/v24Zer0/ToDo/item-service/database"

type ItemHandler struct {
	db *database.Database
}

func NewHandler(db *database.Database) *ItemHandler {
	return &ItemHandler{db: db}
}
