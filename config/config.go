package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/vinitparekh17/project-x/utility"
)

func Init() {
	err := godotenv.Load()
	utility.ErrorHandler(err)
	slog.Info("Config loaded successfully")
}

func GetConfig(env string) (string, error) {
	value := os.Getenv(env)
	if value != "" {
		return value, nil
	} else {
		return "", errors.New(env + " not found in .env file")
	}
}
