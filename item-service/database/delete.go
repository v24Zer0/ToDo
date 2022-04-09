package database

import (
	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB, id string) error {
	item := models.Item{ID: id}

	result := db.Delete(&item)
	return result.Error
}

func DeleteList(db *gorm.DB, id string) error {
	list := models.List{ID: id}

	result := db.Delete(&list)
	return result.Error
}
