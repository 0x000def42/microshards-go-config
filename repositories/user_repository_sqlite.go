package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/0x000def42/microshards-go-config/models"
	"github.com/google/uuid"
)

type UserRepositorySqlite struct {
	db *sql.DB
}

func NewUserRepositorySqlite(db *sql.DB) UserRepository {
	return UserRepositorySqlite{
		db: db,
	}
}

func (repo UserRepositorySqlite) NewUser() *models.User {

	password := uuid.New().String()
	resetToken := uuid.New().String()
	role := models.USER_ROLE_GUEST

	return &models.User{
		Password:   &password,
		ResetToken: &resetToken,
		Role:       &role,
	}
}

func (repo UserRepositorySqlite) GetAll() ([]models.User, error) {
	rows, err := repo.db.Query("select * from users where deleted_at is null")
	if err != nil {
		fmt.Println("[ERROR] UserRepositroy.GetAll: db.query", err)
		return nil, err
	}
	users := []models.User{}
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	return users, nil
}

func (repo UserRepositorySqlite) GetOne(id string) (*models.User, error) {
	getSqlStr := "select * from users where id = $1 and where deleted_at is null"
	row := repo.db.QueryRow(getSqlStr, id)
	user, err := scanUser(row)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo UserRepositorySqlite) Create(user *models.User) (*models.User, error) {
	newId := uuid.New().String()
	createdAt := time.Now()

	user.Id = &newId
	user.CreatedAt = &createdAt

	insertSqlStr := `insert into users 
	(id, username, password, reset_token, role, created_at) 
	values ($1, $2, $3, $4, $5, $6)`

	_, err := repo.db.Exec(insertSqlStr,
		user.Id, user.Username, user.Password, user.ResetToken, user.Role, user.CreatedAt)

	if err != nil {
		return nil, err
	}
	return user, err

}

func (repo UserRepositorySqlite) Update(user *models.User) (*models.User, error) {

	updatedAt := time.Now()

	user.UpdatedAt = &updatedAt

	updateSqlStr := `update users set
		username = $2, password = $3, reset_token = $4, role = $5, updated_at = $6
		where id = $1
	`

	_, err := repo.db.Exec(updateSqlStr,
		user.Id,
		user.Username,
		user.Password,
		user.ResetToken,
		user.Role,
		user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositorySqlite) Delete(user *models.User) error {
	deletedAt := time.Now()
	user.DeletedAt = &deletedAt
	_, err := repo.db.Exec("updates users set deleted_at = $2 where id = $1",
		user.Id,
		user.DeletedAt)

	if err != nil {
		return err
	}
	return nil
}

type ScanableRow interface {
	Scan(dest ...interface{}) error
}

func scanUser(row ScanableRow) (*models.User, error) {
	user := models.User{}
	err := row.Scan(&user.Id,
		&user.Username,
		&user.Password,
		&user.ResetToken,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
