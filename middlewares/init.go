package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/utility"
)

func Init(e *echo.Echo) {
	pwd, _ := os.Getwd()
	file, err := os.OpenFile(pwd+config.K.String("log_path"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utility.ErrorHandler(err)
	defer file.Close()

	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `method=${method}, uri=${uri}, status=${status}, time=${latency_human}`,
		Output:           os.Stdout,
		CustomTimeFormat: "2006-01-02 05:04:05",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.K.Strings("origin"),
	}))

	e.Use(middleware.Secure())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(middleware.RequestID())

	// e.Use(middleware.Recover())
}
