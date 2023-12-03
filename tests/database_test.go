package tests

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestDBConnection(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}

	t.Log("Successfully connected to the database")
}
