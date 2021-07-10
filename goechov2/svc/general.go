package svc

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func BytePassword(password string) []byte {
	return []byte(password)
}
