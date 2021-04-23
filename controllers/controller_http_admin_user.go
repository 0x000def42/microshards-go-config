package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0x000def42/microshards-go-config/app/admin"
	"github.com/0x000def42/microshards-go-config/models"
	"github.com/0x000def42/microshards-go-config/utils"
	"github.com/gorilla/mux"
)

type ControllerHttpAdminUser struct {
	module admin.Module
}

type AdminIndexUserResponse []AdminIndexUserResponsePart

type AdminIndexUserResponsePart struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

func (controller ControllerHttpAdminUser) Index(rw http.ResponseWriter, r *http.Request) {
	users, err := controller.module.UserService.GetList()

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

func (controller ControllerHttpAdminUser) Create(rw http.ResponseWriter, r *http.Request) {

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

	user, err := controller.module.UserService.Create(actionParams)

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

func (controller ControllerHttpAdminUser) Show(rw http.ResponseWriter, r *http.Request) {
	id := getUserID(r)

	user, err := controller.module.UserService.GetOne(id)

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

func (controller ControllerHttpAdminUser) Update(rw http.ResponseWriter, r *http.Request) {
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

	user, err := controller.module.UserService.Update(id, actionParams)

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

func (controller ControllerHttpAdminUser) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getUserID(r)

	err := controller.module.UserService.Delete(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&utils.GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func getUserID(r *http.Request) string {
	vars := mux.Vars(r)
	id := vars["id"]
	return id
}
