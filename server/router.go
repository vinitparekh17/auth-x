package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vinitparekh17/project-x/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())
	v1 := router.Group("v1")
	{
		v1.GET("/health", health.Status2)
	}
	return router

}
