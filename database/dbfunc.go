package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/vinitparekh17/project-x/utility"
)

/*
	This file contains reusable database functions for db operations
	Why we need this?
	To prevent code duplication
*/



func (*db) ModifyData(query string, args ...interface{}) (bool, error) {
	if strings.HasPrefix(query, "INSERT") || strings.HasPrefix(query, "UPDATE") || strings.HasPrefix(query, "DELETE") {
		db := Connect()
		defer Disconnect(db)
		res, err := db.Exec(query, args...)
		utility.ErrorHandler(err)
		return res != nil, err
	}
	return false, errors.New("query must starts with insert, update or delete")
}

func (*db) RetriveData(query string, args ...interface{}) *sql.Row {
	db := Connect()
	defer Disconnect(db)
	row := db.QueryRow(query, args...)
	return row
}
