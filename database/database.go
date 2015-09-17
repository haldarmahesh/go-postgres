package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	_ "log"
	"os"
)

type Config struct {
	adapter string `json:"adapter"`
	host    string `json:"host"`
	dbname  string `json:"dbname"`
	sslmode string `json:"sslmode"`
}

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost dbname=dealscore sslmode=disable")
	logError(err)
	return db
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
