package services

import (
	"log"
	"os"

	"github.com/solnsumei/api-starter-template/models"
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
		log.Fatal(err)
	}

	return db
}

func SyncDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
