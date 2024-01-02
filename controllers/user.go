package controllers

import (
	"errors"
	"net/http"
	"time"

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

	// Binding the request body to the user model
	err := c.Bind(user)
	handler.ErrorHandler(err)

	// Connecting to the database
	db := database.Connect()
	defer database.Disconnect(db)

	// Checking if the email and password is empty
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, utilities.ErrorResponse(utilities.EmptyFIeldErr, errors.New(utilities.LoginFailed)))
	}

	var existEmail string
	er := database.RetriveData(db, "SELECT email FROM \"user\".identity WHERE email = $1", user.Email).Scan(&existEmail)
	handler.ErrorHandler(er)

	// Checking if the user already exists in the database
	if user.Email == existEmail {
		return c.JSON(http.StatusBadRequest, utilities.ErrorResponse("User already exist", errors.New(utilities.SignupFailed)))
	}

	// Hashing the password using bcrypt to securely store it in the database
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	handler.ErrorHandler(err)

	user.Password = string(hashedPass)
	errr := user.Create()
	if errr != nil {
		return c.JSON(http.StatusInternalServerError, errr)
	}
	// Creating a cookie and setting it in the response header
	cookie := new(http.Cookie)
	cookie.Name = "token"
	token := utilities.GenerateJWT(user.Email)
	cookie.Value = token
	cookie.Expires = time.Now().Add(72 * time.Hour)
	c.SetCookie(cookie)

	// Sending the success response
	return c.JSON(http.StatusOK, utilities.SuccessResponse(utilities.SignupSuccess, token))
}

func (*UserControllers) Login(c echo.Context) error {
	user := &models.IdentityModel{}
	// Binding the request body to the user model
	err := c.Bind(&user)
	handler.ErrorHandler(err)

	// Connecting to the database
	db := database.Connect()
	defer database.Disconnect(db)

	// Checking if the email and password is empty
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, utilities.ErrorResponse(utilities.EmptyFIeldErr, errors.New(utilities.LoginFailed)))
	}

	// Checking if the user exists in the database
	type existUser struct {
		Email    string
		Password string
	}
	var eu = &existUser{}

	// Binding the email and password from the database to the existUser struct
	er := database.RetriveData(db, "SELECT email, password FROM \"user\".identity WHERE email = $1", user.Email).Scan(&eu.Email, &eu.Password)
	handler.ErrorHandler(er)

	// Comparing the password from the database with the password from the request body
	if user.Email == eu.Email {
		err := bcrypt.CompareHashAndPassword([]byte(eu.Password), []byte(user.Password))
		if err == nil {
			cookie := new(http.Cookie)
			cookie.Name = "token"
			token := utilities.GenerateJWT(user.Email)
			cookie.Value = token
			cookie.Expires = time.Now().Add(72 * time.Hour)
			c.SetCookie(cookie)
			return c.JSON(http.StatusOK, utilities.SuccessResponse(utilities.LoginSuccess, token))
		}
		// Sending the error response if the password is incorrect
		return c.JSON(http.StatusUnauthorized, utilities.ErrorResponse("Incorrect password, try again!", errors.New(utilities.LoginFailed)))
	}
	// Sending the error response if the user does not exists
	return c.JSON(http.StatusUnauthorized, utilities.ErrorResponse("User with this email does not exists", errors.New(utilities.LoginFailed)))
}

func (*UserControllers) Logout(c echo.Context) error {
	// Checking if the user is logged in or not
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utilities.ErrorResponse("User is not logged in", errors.New(utilities.LogoutFailed)))
	}

	// Deleting the cookie from the response header to logout the user
	cookie.Expires = time.Now()
	c.SetCookie(cookie)

	// Sending the success response
	return c.JSON(http.StatusOK, utilities.SuccessResponse("User has been logged out successfully", nil))
}
