package config

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/vinitparekh17/project-x/handler"
)

var K *koanf.Koanf

func LoadConfig() {
	K = koanf.New(".")
	err := K.Load(file.Provider("config.yml"), yaml.Parser())
	handler.ErrorHandler(err)
	slog.Info("Config. loaded successfully")
}

func LoadEnv() {
	err := godotenv.Load()
	handler.ErrorHandler(err)
	slog.Info("Env. loaded successfully")
}

func GetConfig(env string) (string, error) {
	value := os.Getenv(env)
	if value != "" {
		return value, nil
	} else {
		return "", errors.New(env + " not found in .env file")
	}
}
