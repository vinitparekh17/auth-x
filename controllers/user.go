package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/database"
	"github.com/vinitparekh17/project-x/handler"
	"github.com/vinitparekh17/project-x/models"
	"github.com/vinitparekh17/project-x/utilities"
	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct{}

func (*UserControllers) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func (*UserControllers) Signup(c echo.Context) error {
	user := &models.IdentityModel{}

	err := c.Bind(user)
	handler.ErrorHandler(err)

	db := database.Connect()
	defer database.Disconnect(db)

	var existEmail string
	row := database.RetriveData(db, "SELECT email FROM \"user\".identity WHERE email = $1", user.Email)
	for row.Next() {
		err := row.Scan(&existEmail)
		handler.ErrorHandler(err)
	}
	if user.Email != existEmail {
		if user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, "Email and Password are required")
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		handler.ErrorHandler(err)

		user.Password = string(hashedPass)
		errr := user.Create(*user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errr)
		}
		return c.JSON(http.StatusOK, utilities.SuccessResponse(utilities.SignupSuccess, user))
	}

	return c.JSON(http.StatusBadRequest, utilities.ErrorResponse("User already exist", errors.New(utilities.SignupFailed)))
}

func (*UserControllers) Login(c echo.Context) error {
	user := &models.IdentityModel{}

	err := c.Bind(&user)
	handler.ErrorHandler(err)

	db := database.Connect()
	defer database.Disconnect(db)

	type existUser struct {
		Email    string
		Password string
	}
	var eu = &existUser{}
	row := database.RetriveData(db, "SELECT email, password FROM \"user\".identity WHERE email = $1", user.Email)
	for row.Next() {
		err := row.Scan(&eu.Email, &eu.Password)
		handler.ErrorHandler(err)
	}
	if user.Email == eu.Email {
		err := bcrypt.CompareHashAndPassword([]byte(eu.Password), []byte(user.Password))
		if err == nil {
			return c.JSON(http.StatusOK, utilities.SuccessResponse(utilities.LoginSuccess, eu.Email))
		}
		return c.JSON(http.StatusUnauthorized, utilities.ErrorResponse("Password has been incorrect", errors.New(utilities.LoginFailed)))
	}
	return c.JSON(http.StatusNotFound, utilities.ErrorResponse("User with this email does not exists", errors.New(utilities.LoginFailed)))
}
