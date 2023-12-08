package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vinitparekh17/project-x/utility"
)

func Init(e *echo.Echo) {
	pwd, _ := os.Getwd()
	file, err := os.OpenFile(pwd+"/logss/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utility.ErrorHandler(err)
	defer file.Close()
	e.Use(middleware.BodyLimit("2M"))
		
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
		Output: file,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
}
