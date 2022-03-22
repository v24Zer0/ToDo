package database

import (
	"github.com/segmentio/ksuid"
	"github.com/v24Zer0/ToDo/item-service/models"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) {
	id := ksuid.New()

	item.ID = id.String()
	db.Create(item)
}

func CreateList(db *gorm.DB, list *models.List) {
	id := ksuid.New()

	list.ID = id.String()

	db.Create(list)
}
