package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func Init() {
	path := filepath.Join("database", "user.sql")
	reader, err := os.ReadFile(path)
	handler.ErrorHandler(err)
	queries := string(reader)
	pg := Connect()
	defer Disconnect(pg)
	_, execerr := pg.Exec(queries)
	handler.ErrorHandler(execerr)
}
func Connect() *sql.DB {
	connStr, err := config.GetConfig("POSTGRES_URL")
	handler.ErrorHandler(err)
	pg, err := sql.Open("postgres", connStr)
	handler.ErrorHandler(err)
	slog.Info("Database connected successfully")
	return pg
}

func Disconnect(db *sql.DB) {
	err := db.Close()
	handler.ErrorHandler(err)
	slog.Info("Database closed successfully")
}
