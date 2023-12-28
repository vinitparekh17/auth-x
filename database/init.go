package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
)

func Init() {
	pg := Connect()
	defer Disconnect(pg)
	er := pg.Ping()
	handler.ErrorHandler(er)
	if er == nil {
		slog.Info("Database pinged successfully")
	}
	// path := filepath.Join("database", "user.sql")
	// reader, err := os.ReadFile(path)
	// handler.ErrorHandler(err)
	// queries := string(reader)
	// pg.Exec(queries)
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
