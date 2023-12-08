package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/vinitparekh17/project-x/utility"
)

var k = koanf.New(".")

func LoadConfig() {
	err := k.Load(file.Provider("config.yml"), yaml.Parser())
	utility.ErrorHandler(err)
}

func LoadEnv() {
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
