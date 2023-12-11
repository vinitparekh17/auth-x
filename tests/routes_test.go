package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/vinitparekh17/project-x/controllers"
	"github.com/vinitparekh17/project-x/handler"
)

func TestGetHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health/get", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	assert.Equal(t, http.StatusOK, rec.Code)
	c := e.NewContext(req, rec)
	h := (&controllers.HealthController{})
	host, err := os.Hostname()
	handler.ErrorHandler(err)
	if assert.NoError(t, h.GetHealth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\""+host+" is healthy\"\n", rec.Body.String())
	}
}
