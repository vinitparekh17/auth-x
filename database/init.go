package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	_ "github.com/lib/pq"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/utility"
)

func Init() {
	path := filepath.Join("database", "user.sql")
	reader, err := os.ReadFile(path)
	utility.ErrorHandler(err)
	queries := string(reader)
	pg := Connect()
	defer Disconnect(pg)
	_, execerr := pg.Exec(queries)
	utility.ErrorHandler(execerr)
}

func Connect() *sql.DB {
	connStr, err := config.GetConfig("POSTGRES_URL")
	utility.ErrorHandler(err)
	pg, err := sql.Open("postgres", connStr)
	utility.ErrorHandler(err)
	slog.Info("Database connected successfully")
	return pg
}

func Disconnect(db *sql.DB) {
	err := db.Close()
	utility.ErrorHandler(err)
	slog.Info("Database closed successfully")
}
