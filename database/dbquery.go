package database

import (
	"fmt"
	"strings"
)

/*
	This file contains the structure of database queries
	The Question: Why we need this?
	To prevent SQL Injection, better code readability and maintainability and to prevent the code from becoming a mess
	Also, we can use this to generate the queries dynamically as per my perseption
*/

type Insert struct {
	Table  string
	Fields []string
}

type Update struct {
	Table  string
	Fields []string
	Where  string
}

type Delete struct {
	Table string
	Where string
}

type All struct {
	Table string
}

func (i *Insert) Build() string {
	placeholders := make([]string, len(i.Fields))
	for idx := range i.Fields {
		placeholders[idx] = "$" + fmt.Sprintf("%d", idx+1)
	}
	return `INSERT INTO ` + i.Table + ` (` + strings.Join(i.Fields, ", ") + `) VALUES (` + strings.Join(placeholders, ", ") + `)`
}

func (u *Update) Build() string {
	placeholders := make([]string, len(u.Fields))
	for idx := range u.Fields {
		placeholders[idx] = "$" + fmt.Sprintf("%d", idx+1)
	}
	return `UPDATE ` + u.Table + ` SET ` + strings.Join(u.Fields, ",") + `=` + strings.Join(placeholders, ", ") + ` WHERE ID=` + `$` + fmt.Sprintf("%d", len(u.Fields)+1)
}

func (d *Delete) Build() string {
	return `DELETE FROM ` + d.Table + `WHERE ID = $1`
}

func (a *All) Build() string {
	return `SELECT * FROM ` + a.Table
}
