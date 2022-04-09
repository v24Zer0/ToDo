package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
}
