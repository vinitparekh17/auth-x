package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/vinitparekh17/project-x/controllers"
)

type UserController struct {
	*controllers.UserControllers
}

func NewRouter() chi.Router {
	r := chi.NewRouter()
	return r
}

func (r UserController) Routes() chi.Router {
	router := NewRouter()
	router.Get("/", r.GetUser)
	return router
}
