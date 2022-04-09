package database

import (
	"errors"

	"github.com/v24Zer0/ToDO/user-service/models"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, id string) error {
	var u models.User
	db.Where(models.User{ID: id}).First(&u)
	if u.ID == "" {
		return errors.New("user does not exist")
	}

	result := db.Delete(&u)
	return result.Error
}
