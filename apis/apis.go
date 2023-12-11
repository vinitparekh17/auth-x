package apis

import (
	"github.com/labstack/echo/v4"
	"github.com/vinitparekh17/project-x/controllers"
)

func Init(e *echo.Echo) {
	user := e.Group("/user")
	user.POST("/signup", (&controllers.UserControllers{}).Signup)
	user.POST("/login", (&controllers.UserControllers{}).Login)

	health := e.Group("/health")
	health.GET("/get", (&controllers.HealthController{}).GetHealth)
}
