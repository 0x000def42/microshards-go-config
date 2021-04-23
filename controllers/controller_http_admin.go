package controllers

import (
	"fmt"
	"net/http"

	"github.com/0x000def42/microshards-go-config/app/admin"
	"github.com/gorilla/mux"
)

type ControllerHttpAdmin struct {
	module admin.Module
	User   ControllerHttpAdminUser
}

func NewControllerHttpAdmin(adminModule admin.Module) ControllerHttp {
	return &ControllerHttpAdmin{
		module: adminModule,
		User:   ControllerHttpAdminUser{module: adminModule},
	}
}

func (controller ControllerHttpAdmin) Routes(router *mux.Router) {
	admin := router.PathPrefix("/admin/").Subrouter()
	admin.Use(controller.AdminMiddleware)

	id := "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	userPath := fmt.Sprintf("/users/{id:%s}", id)
	admin.HandleFunc("/users", controller.User.Index).Methods(http.MethodGet)
	admin.HandleFunc("/users", controller.User.Create).Methods(http.MethodPatch)
	admin.HandleFunc(userPath, controller.User.Show).Methods(http.MethodGet)
	admin.HandleFunc(userPath, controller.User.Update).Methods(http.MethodPatch)
	admin.HandleFunc(userPath, controller.User.Update).Methods(http.MethodPut)
	admin.HandleFunc(userPath, controller.User.Delete).Methods(http.MethodDelete)
}

func (handler ControllerHttpAdmin) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
