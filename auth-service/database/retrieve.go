package database

import (
	"errors"

	"github.com/v24Zer0/ToDo/auth-service/models"
	"gorm.io/gorm"
)

func RetrieveToken(db *gorm.DB, userID string) (*models.UserToken, error) {
	var t models.UserToken
	db.Where(models.UserToken{UserID: userID}).First(&t)
	if t.Token == "" {
		return nil, errors.New("token not found")
	}

	return &t, nil
}

func RetrieveByToken(db *gorm.DB, token string) (*models.UserToken, error) {
	var t models.UserToken
	db.Where(models.UserToken{Token: token}).First(&t)
	if t.UserID == "" {
		return nil, errors.New("token not found")
	}

	return &t, nil
}
