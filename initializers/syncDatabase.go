package initializers

import "github.com/RakibSiddiquee/gin-jwt-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
