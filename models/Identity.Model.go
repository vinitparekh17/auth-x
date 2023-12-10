package models

import (
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/utility"
)

/*
	Identity model that describes how our Identity table looks like
	This table is contains auth specific data i.e. email and password hash
	I did like to keep it minimal and simple to avoid any security issues
*/

type IdentityModel struct {
	UID      int64  `json:"id,omitempty" unique:"true"`
	Email    string `json:"email" unique:"true"`
	Password string `json:"password" min:"8"`
}

func (*IdentityModel) Create(usr IdentityModel) error {
	db := database.Connect()
	defer database.Disconnect(db)
	query := database.Insert{
		Table:  "identity",
		Fields: []string{"email", "password"},
		Values: []interface{}{usr.Email, usr.Password},
	}
	_, err := db.Exec(query.Build(), query.Values...)
	utility.ErrorHandler(err)
	return err
}
