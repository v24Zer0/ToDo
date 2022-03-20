package handlers

import (
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db: db}
}
