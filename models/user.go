package models

import "time"

type User struct {
	Id         *string
	Username   *string
	Password   *string
	ResetToken *string
	Role       *UserRole
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}

type UserRole int

// Define consts for value of RoleType field
const (
	USER_ROLE_GUEST UserRole = 1
	USER_ROLE_USER  UserRole = 2
	USER_ROLE_ADMIN UserRole = 3
)
