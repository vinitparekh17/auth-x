package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/models"
	"github.com/vinitparekh17/project-x/utility"
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

	return c.JSON(http.StatusBadRequest, "Email and Password are required")
}