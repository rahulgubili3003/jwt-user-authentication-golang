package init

import (
	"github.com/joho/godotenv"
	"log"
)

func InitializeEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}
