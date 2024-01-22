package common

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	err := godotenv.Load()
	if err != nil {
		helpers.ErrorLogger.Panicln("Error loading.env file")
	}

	return os.Getenv("MONGO_URI")
}
