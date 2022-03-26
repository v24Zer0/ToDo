package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	dbURL := os.Getenv("dbURL")
	dbName := os.Getenv("dbName")
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")

	dbPath := fmt.Sprintf("%s:%s@tcp%s/%s", dbUser, dbPass, dbURL, dbName)

	newLogger := logger.New(log.New(os.Stdout, "[gorm]", log.LstdFlags), logger.Config{
		SlowThreshold:             time.Second,
		Colorful:                  false,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Silent,
	})

	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
