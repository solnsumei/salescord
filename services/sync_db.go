package services

import (
	"github.com/solnsumei/api-starter-template/models"
	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
