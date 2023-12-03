package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserControllers struct{}

func (rs UserControllers) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.GetUser)
	return r
}

func (*UserControllers) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
