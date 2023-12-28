package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/vinitparekh17/project-x/handler"
)

/*
	This file contains reusable database functions for db operations
	Why we need this?
	To prevent code duplication
*/

func ModifyData(query string, args ...interface{}) (bool, error) {
	if strings.HasPrefix(query, "INSERT") || strings.HasPrefix(query, "UPDATE") || strings.HasPrefix(query, "DELETE") {
		db := Connect()
		defer Disconnect(db)
		res, err := db.Exec(query, args...)
		handler.ErrorHandler(err)
		return res != nil, err
	}
	return false, errors.New("query must starts with insert, update or delete")
}

func RetriveData(db *sql.DB, query string, args ...interface{}) *sql.Row {
	row := db.QueryRow(query, args...)
	return row
}
