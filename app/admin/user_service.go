package admin

import (
	"fmt"

	"github.com/0x000def42/microshards-go-config/models"
	"github.com/0x000def42/microshards-go-config/repositories"
)

type IUserService interface {
	GetList() ([]models.User, error)
	Create(CreateUserParams) (*models.User, error)
	Update(id string, params UpdateUserParams) (*models.User, error)
	GetOne(id string) (*models.User, error)
	Delete(id string) error
}

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) IUserService {
	return UserService{
		repo: repo,
	}
}

func (service UserService) GetList() ([]models.User, error) {
	users, err := service.repo.GetAll()

	if err != nil {
		fmt.Println("[ERROR] admin.UserService repo.GetAll", err)
		return nil, err
	}

	return users, nil
}

type CreateUserParams struct {
	Username string          `validate:"required"`
	Role     models.UserRole `validate:"required,user_role"`
}

func (service UserService) Create(params CreateUserParams) (*models.User, error) {
	user := service.repo.NewUser()

	user.Role = &params.Role
	user.Username = &params.Username

	user, err := service.repo.Create(user)

	if err != nil {
		return nil, err
	}

	// service.event_store.PublishUserCreated()

	return user, nil

}

type UpdateUserParams struct {
	Username *string
	Role     *models.UserRole
}

func (service UserService) Update(id string, params UpdateUserParams) (*models.User, error) {
	user, err := service.repo.GetOne(id)

	if err != nil {
		return nil, err
	}

	if params.Role != nil {
		user.Role = params.Role
	}
	if params.Username != nil {
		user.Username = params.Username
	}

	user, err = service.repo.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service UserService) GetOne(id string) (*models.User, error) {
	user, err := service.repo.GetOne(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service UserService) Delete(id string) error {
	user, err := service.repo.GetOne(id)

	if err != nil {
		return err
	}

	err = service.repo.Delete(user)

	if err != nil {
		return err
	}

	return nil
}
