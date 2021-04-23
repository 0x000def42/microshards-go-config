package controllers

import "github.com/gorilla/mux"

type ControllerHttp interface {
	Routes(*mux.Router)
}
