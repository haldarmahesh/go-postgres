package database

import (
	"database/sql"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost dbname=delscore sslmode=disable")
	logError(err)
	return db
}

func logError(err error) {
  if err != nil {
    panic(err)
  }
}