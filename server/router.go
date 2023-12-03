package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinitparekh17/project-x/controllers"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Mount("/health", controllers.HealthController{}.Routes())
	r.Mount("/user", controllers.UserControllers{}.Routes())
	return r
}
