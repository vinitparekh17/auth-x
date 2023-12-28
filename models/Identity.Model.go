package models

import (
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/handler"
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

type PublicCredential struct {
	Email string `json:"email"`
}

func (u *IdentityModel) Create() error {
	db := database.Connect()
	defer database.Disconnect(db)
	query := database.Insert{
		Table:  `"user".identity`,
		Fields: []string{"email", "password"},
	}

	smt, err := db.Prepare(query.Build())
	handler.ErrorHandler(err)
	defer smt.Close()
	_, er := smt.Exec(u.Email, u.Password)
	handler.ErrorHandler(er)
	return nil
}
