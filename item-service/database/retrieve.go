package database

import (
	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func RetrieveItems(db *gorm.DB, listID string) *models.Items {
	var items models.Items
	db.Where(&models.Item{ListID: listID}).Find(&items)
	return &items
}

func RetrieveLists(db *gorm.DB, userID string) *models.Lists {
	var lists models.Lists
	db.Where(&models.List{UserID: userID}).Find(&lists)
	return &lists
}
