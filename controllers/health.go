package controllers

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type HealthController struct{}

func (rs HealthController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.GetHealth)
	return r
}

func (*HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Write([]byte("Hostname: " + hostname + "is healthy"))
}
