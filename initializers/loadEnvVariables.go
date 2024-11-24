package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("DB_URL") == "" {
		log.Fatal("DB_URL environment variable is required but not set.")
	}
}
