package admin

import (
	"github.com/0x000def42/microshards-go-config/models"
)

type CreateUserParams struct {
	Username string
	Role     models.UserRole
}

type UpdateUserParams struct {
	Username *string
	Role     *models.UserRole
}
