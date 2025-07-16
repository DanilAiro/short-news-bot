package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		Log.Fatal("Error loading .env file")
	}
}