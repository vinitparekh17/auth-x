package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/utility"
)

func Init() *sql.DB {
	connStr, err := config.GetConfig("POSTGRES_URL")
	utility.ErrorHandler(err)
	db, err := sql.Open("postgres", connStr)
	utility.ErrorHandler(err)
	slog.Info("Database connected successfully")
	err = db.Ping()
	utility.ErrorHandler(err)
	slog.Info("Database pinged successfully")
	CreateDb(db)

	return db
}

func CreateDb(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS Project-x")
	utility.ErrorHandler(err)
	slog.Info("Database created successfully")
}
