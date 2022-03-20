package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	dbURL := os.Getenv("dbURL")
	dbName := os.Getenv("dbName")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")

	dbPath := fmt.Sprintf("%s:%s@tcp%s/%s", dbUser, dbPass, dbURL, dbName)

	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}
