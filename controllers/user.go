package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/models"
	"github.com/vinitparekh17/project-x/utility"
	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct{}

func (*UserControllers) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func (*UserControllers) Signup(c echo.Context) error {
	user := &models.IdentityModel{}
	err := c.Bind(user)
	utility.ErrorHandler(err)

	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, "Email and Password are required")
	}

	db := database.Connect()
	defer database.Disconnect(db)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	utility.ErrorHandler(err)

	user.Password = string(hashedPass)
	errr := user.Create(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errr)
	}
	return c.JSON(http.StatusOK, user)
}

func (*UserControllers) Login(c echo.Context) error {
	user := &models.UserModel{}
	db := database.Connect()
	defer database.Disconnect(db)
	err := c.Bind(&user)
	utility.ErrorHandler(err)
	// To be continued...
	return nil
}
