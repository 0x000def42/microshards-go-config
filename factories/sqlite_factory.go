package factories

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteClient(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println("[ERROR] Can't open sqlite database", err)
		panic(err)
	}
	return db
}
