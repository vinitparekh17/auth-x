package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllers struct{}

func (*UserControllers) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
