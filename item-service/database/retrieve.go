package database

import (
	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func RetrieveItems(db *gorm.DB, listID string) (*models.Items, error) {
	var items models.Items
	result := db.Where(&models.Item{ListID: listID}).Find(&items)
	return &items, result.Error
}

func RetrieveLists(db *gorm.DB, userID string) (*models.Lists, error) {
	var lists models.Lists
	result := db.Where(&models.List{UserID: userID}).Find(&lists)
	return &lists, result.Error
}
