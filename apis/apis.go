package apis

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/vinitparekh17/project-x/controllers"
	_ "github.com/vinitparekh17/project-x/docs"
)

func Init(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	user := e.Group("/user")
	user.POST("/signup", (&controllers.UserControllers{}).Signup)
	user.POST("/login", (&controllers.UserControllers{}).Login)

	health := e.Group("/health")
	health.GET("/get", (&controllers.HealthController{}).GetHealth)
}
