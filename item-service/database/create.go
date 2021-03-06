package database

import (
	"errors"

	"github.com/segmentio/ksuid"
	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) error {
	id := ksuid.New()

	item.ID = id.String()
	result := db.Create(item)
	return result.Error
}

func CreateList(db *gorm.DB, list *models.List) error {
	var l models.List

	db.First(&l, list)
	if l.ID != "" {
		return errors.New("item already exists")
	}

	id := ksuid.New()

	list.ID = id.String()
	result := db.Create(list)
	return result.Error
}
