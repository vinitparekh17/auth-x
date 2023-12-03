package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	os.Exit(m.Run())
}

func TestEnv(t *testing.T) {
	t.Log("Testing environment variables")
	// TODO: Add tests for environment variables
	POSTGRES_URL := os.Getenv("POSTGRES_URL")
	if POSTGRES_URL == "" {
		t.Error("POSTGRES_URL not found in .env file")
	}
}
