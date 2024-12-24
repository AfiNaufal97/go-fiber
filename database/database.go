package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() {
	connStr := "postgres://postgres:afi@localhost:5432/books_api?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic("ada erorr")
	}

	DB = db

}
