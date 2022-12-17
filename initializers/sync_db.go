package initializers

import "github.com/solnsumei/api-starter-template/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
