package request_handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0x000def42/microshards-go-config/app/admin"
	"github.com/0x000def42/microshards-go-config/models"
	"github.com/0x000def42/microshards-go-config/utils"
	"github.com/gorilla/mux"
)

type RequestHandlerHttpAdmin struct {
	module admin.Module
}

func NewRequestHandlerHttpAdmin(adminModule admin.Module) RequestHandlerHttp {
	return &RequestHandlerHttpAdmin{
		module: adminModule,
	}
}

func (handler RequestHandlerHttpAdmin) Routes(router *mux.Router) {
	admin := router.PathPrefix("/admin/").Subrouter()
	admin.Use(handler.AdminMiddleware)

	id := "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	userPath := fmt.Sprintf("/users/{id:%s}", id)
	admin.HandleFunc("/users", handler.adminIndexUser).Methods("GET")
	admin.HandleFunc("/users", handler.adminCreateUser).Methods("POST")
	admin.HandleFunc(userPath, handler.adminShowUser).Methods("GET")
	admin.HandleFunc(userPath, handler.adminUpdateUser).Methods("PATCH")
}

func (handler RequestHandlerHttpAdmin) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

type AdminIndexUserResponse []AdminIndexUserResponsePart

type AdminIndexUserResponsePart struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

func (handler RequestHandlerHttpAdmin) adminIndexUser(rw http.ResponseWriter, r *http.Request) {
	users, err := handler.module.UserService.GetList()

	if err != nil {
		fmt.Println("[ERROR] GET /admin/users", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := AdminIndexUserResponse{}

	for _, user := range users {
		response = append(response, AdminIndexUserResponsePart{
			Id:       *user.Id,
			Username: *user.Username,
			Role:     *user.Role,
		})
	}

	err = utils.ToJSON(response, rw)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

type AdminCreateUserParams struct {
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

type AdminCreateUserResponse struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

func (handler RequestHandlerHttpAdmin) adminCreateUser(rw http.ResponseWriter, r *http.Request) {

	params := &AdminCreateUserParams{}
	err := utils.FromJSON(params, r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	actionParams := admin.CreateUserParams{
		Username: params.Username,
		Role:     params.Role,
	}

	errs := utils.Validate(actionParams)

	if len(errs) != 0 {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		utils.ToJSON(&utils.ValidationErrorMessages{Messages: errs.Errors()}, rw)
		return
	}

	user, err := handler.module.UserService.Create(actionParams)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	response := AdminCreateUserResponse{
		Id:       *user.Id,
		Username: *user.Username,
		Role:     *user.Role,
	}

	err = utils.ToJSON(response, rw)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)

}

type AdminShowUserResponse struct {
	Id         string          `json:"id"`
	Username   string          `json:"username"`
	Role       models.UserRole `json:"role"`
	ResetToken string          `json:"reset_token"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  *time.Time      `json:"updated_at"`
}

func (handler RequestHandlerHttpAdmin) adminShowUser(rw http.ResponseWriter, r *http.Request) {
	id := getUserID(r)

	user, err := handler.module.UserService.GetOne(id)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&utils.GenericError{Message: "Not found"}, rw)
		return
	}

	response := AdminShowUserResponse{
		Id:         *user.Id,
		Username:   *user.Username,
		Role:       *user.Role,
		ResetToken: *user.ResetToken,
		CreatedAt:  *user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	err = utils.ToJSON(response, rw)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

type AdminUpdateUserParams struct {
	Username *string          `json:"username"`
	Role     *models.UserRole `json:"role"`
}

type AdminUpdateUserResponse struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

func (handler RequestHandlerHttpAdmin) adminUpdateUser(rw http.ResponseWriter, r *http.Request) {
	id := getUserID(r)
	params := &AdminUpdateUserParams{}
	err := utils.FromJSON(params, r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	actionParams := admin.UpdateUserParams{
		Username: params.Username,
		Role:     params.Role,
	}

	errs := utils.Validate(actionParams)

	if len(errs) != 0 {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		utils.ToJSON(&utils.ValidationErrorMessages{Messages: errs.Errors()}, rw)
		return
	}

	user, err := handler.module.UserService.Update(id, actionParams)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	response := AdminUpdateUserResponse{
		Id:       *user.Id,
		Username: *user.Username,
		Role:     *user.Role,
	}

	err = utils.ToJSON(response, rw)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func getUserID(r *http.Request) string {
	vars := mux.Vars(r)
	id := vars["id"]
	return id
}
