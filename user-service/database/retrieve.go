package database

import (
	"errors"

	"github.com/v24Zer0/ToDO/user-service/models"
	"gorm.io/gorm"
)

func RetrieveUser(db *gorm.DB, user *models.User) (*models.User, error) {
	var u models.User
	db.Where(models.User{Username: user.Username}).First(&u)
	if u.ID == "" {
		return nil, errors.New("user does not exist")
	}

	return &u, nil
}
