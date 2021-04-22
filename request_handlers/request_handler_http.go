package request_handlers

import "github.com/gorilla/mux"

type RequestHandlerHttp interface {
	Routes(*mux.Router)
}
