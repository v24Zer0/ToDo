package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dbURL := os.Getenv("dbURL")
	dbName := os.Getenv("dbName")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")

	dbPath := fmt.Sprintf("%s:%s@tcp%s/%s", dbUser, dbPass, dbURL, dbName)

	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
