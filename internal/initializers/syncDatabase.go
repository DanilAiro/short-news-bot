package initializers

import "short-news-bot/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}