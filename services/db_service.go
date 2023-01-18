package services

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	var err error
	var db *gorm.DB

	if os.Getenv("DB_TYPE") == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to DB")
	}

	return db
}
