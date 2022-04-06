package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/v24Zer0/ToDo/auth-service/models"
	"gorm.io/gorm"
)

func CreateToken(db *gorm.DB, userID string) error {
	var t models.UserToken
	db.Where(models.UserToken{UserID: userID}).First(&t)
	if t.Token != "" {
		return errors.New("user already has token")
	}

	randomToken := uuid.NewString()
	t.UserID = userID
	t.Token = randomToken

	result := db.Create(&t)
	return result.Error
}
