package apis

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/controllers"
)

func Init(e *echo.Echo) {
	user := e.Group("/user")
	user.GET("/get", (&controllers.UserControllers{}).GetUsers)

	health := e.Group("/health")
	health.GET("/get", (&controllers.HealthController{}).GetHealth)
}
