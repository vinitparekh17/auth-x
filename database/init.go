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
	return db
}
