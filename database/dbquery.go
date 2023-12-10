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
	Values []interface{}
}

type Update struct {
	Table  string
	Fields []string
	Values []interface{}
	Where  string
}

type Delete struct {
	Table string
	Where string
}

type Select struct {
	Table  string
	Fields []string
	Where  string
}

type All struct {
	Table string
}

func (i *Insert) Build() string {
	placeholders := make([]string, len(i.Fields))
	values := make([]interface{}, len(i.Fields))
	for idx, _ := range i.Fields {
		placeholders[idx] = "$" + fmt.Sprintf("%d", idx+1)
		values[idx] = i.Values[idx]
	}
	return `INSERT INTO ` + i.Table + ` (` + strings.Join(i.Fields, ",") + `) VALUES (` + strings.Join(placeholders, ",") + `)`
}

func (u *Update) Build() string {
	// query := `UPDATE ` + u.Table + ` SET ` + strings.Join(u.Fields, ",") + `=` + strings.Join(u.Values, ",") + ` WHERE ` + u.Where
	return "query"
}
