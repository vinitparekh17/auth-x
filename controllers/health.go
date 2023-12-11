package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/handler"
)

type HealthController struct{}

func (*HealthController) GetHealth(c echo.Context) error {
	hostname, err := os.Hostname()
	handler.ErrorHandler(err)
	return c.JSON(http.StatusOK, hostname+" is healthy")
}
