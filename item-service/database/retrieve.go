package database

import "github.com/v24Zer0/ToDo/item-service/models"

func (db *Database) RetrieveItems(listID string) *models.Items {
	var items models.Items
	db.db.Where(&models.Item{ListID: listID}).Find(&items)
	return &items
}

func (db *Database) RetrieveLists(userID string) *models.Lists {
	var lists models.Lists
	db.db.Where(&models.List{UserID: userID}).Find(&lists)
	return &lists
}
