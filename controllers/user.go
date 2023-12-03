package controllers

import (
	"net/http"
)

type UserControllers struct{}

func (*UserControllers) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
