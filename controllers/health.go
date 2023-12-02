package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (*HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (*HealthController) Status2(c *gin.Context) {
	c.String(http.StatusOK, "V1 Working!")
}
