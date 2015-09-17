package database

import (
	"database/sql"
	_ "encoding/json"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "os"
)

type Config struct {
	adapter string `json:"adapter"`
	host    string `json:"host"`
	dbname  string `json:"dbname"`
	sslmode string `json:"sslmode"`
}

func CreateConnection() *sql.DB {
	/*content, err := io.ioutil.ReadFile("config.json")
	logError(err)
	fmt.Println(string(content))*/
	db, err := sql.Open("postgres", "host=localhost dbname=dealscore sslmode=disable")
	logError(err)
	return db
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
