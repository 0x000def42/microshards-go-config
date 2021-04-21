package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct{}

func NewClient(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
