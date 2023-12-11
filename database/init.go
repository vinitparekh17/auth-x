package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func Connect() *sql.DB {
	connStr, err := config.GetConfig("POSTGRES_URL")
	utility.ErrorHandler(err)
	db, err := sql.Open("postgres", connStr)
	utility.ErrorHandler(err)
	slog.Info("Database connected successfully")
	CreateDb(db)
	return db
}

func Disconnect(db *sql.DB) {
	err := db.Close()
	handler.ErrorHandler(err)
	slog.Info("Database closed successfully")
}

func CreateDb(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS Project-x")
	utility.ErrorHandler(err)
	slog.Info("Database created successfully")
}
