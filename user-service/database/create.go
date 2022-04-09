package database

import (
	"errors"

	"github.com/segmentio/ksuid"
	"github.com/v24Zer0/ToDO/user-service/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	var u models.User

	db.Where(models.User{Username: user.Username}).First(&u)
	if u.ID != "" {
		return errors.New("username already exists")
	}

	hashpwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	id := ksuid.New()
	user.ID = id.String()
	user.Password = string(hashpwd)

	result := db.Create(&user)
	return result.Error
}
