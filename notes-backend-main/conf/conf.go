package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("error: error loading .env file")
	}
	return os.Getenv(key)
}
