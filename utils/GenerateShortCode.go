package utils

import (
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"
)

func GenerateShortCode() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	charPool := os.Getenv("CHAR_POOL")
	const LENGTH = 6
	var shortCode string

	for range LENGTH {
		randomIndex := rand.Intn(len(charPool))
		shortCode = shortCode + string(charPool[randomIndex])
	}

	return shortCode
}
