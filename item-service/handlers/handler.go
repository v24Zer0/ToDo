package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db: db}
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
}
