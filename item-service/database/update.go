package database

import (
	"errors"

	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB, item *models.Item) error {
	var i models.Item

	db.Where(&models.Item{ID: item.ID}).First(&i)
	if i.ID == "" {
		return errors.New("item does not exist")
	}

	item.ListID = i.ListID
	result := db.Save(item)
	return result.Error
}

func UpdateList(db *gorm.DB, list *models.List) error {
	var l models.List

	db.Where(&models.List{ID: list.ID}).First(&l)
	if l.ID == "" {
		return errors.New("list does not exist")
	}

	list.UserID = l.UserID
	result := db.Save(list)
	return result.Error
}
