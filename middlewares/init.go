package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
