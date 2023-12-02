package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetConfig(env string) (string, error) {
	value := os.Getenv(env)
	if value != "" {
		return value, nil
	} else {
		return "", errors.New(env + " not found in .env file")
	}
}
