package repositories

import "github.com/0x000def42/microshards-go-config/models"

type UserRepository interface {
	NewUser() *models.User
	GetAll() ([]models.User, error)
	GetOne(id string) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	Delete(*models.User) error
}
