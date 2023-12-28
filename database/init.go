package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func Init() {
	// path := filepath.Join("database", "user.sql")
	// reader, err := os.ReadFile(path)
	// handler.ErrorHandler(err)
	// queries := string(reader)
	pg := Connect()
	defer Disconnect(pg)
	err := pg.Ping()
	handler.ErrorHandler(err)
	if err == nil {
		slog.Info("Database pinged successfully")
	}
}
func Connect() *sql.DB {
	connStr, err := config.GetEnv("POSTGRES_URL")
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
