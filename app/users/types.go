package app

// Define user entity
type User struct {
	Id         string
	Username   string
	Password   string
	ResetToken string
	Role       UserRole
}

// Define subtype for Role field in User type
type UserRole int

// Define consts for value of RoleType field
const (
	USER_ROLE_GUEST UserRole = 1
	USER_ROLE_USER  UserRole = 2
	USER_ROLE_ADMIN UserRole = 3
)

// Define user repository struct
type UserRepository interface {
	All() ([]User, error)
	Save(user User) (User, error)
	Get(id int) (User, error)
}
