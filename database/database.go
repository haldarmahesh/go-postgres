package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "log"
	_ "os"
	"path"
	"path/filepath"
	"runtime"
)

type Config struct {
	Adapter string `json:"adapter"`
	Host    string `json:"host"`
	Dbname  string `json:"dbname"`
	Sslmode string `json:"sslmode"`
}

func CreateConnection() *sql.DB {
	data := readJsonConfig()
	configuration := Config{}
	err := json.Unmarshal(data, &configuration)
	logError(err)
	db, err := sql.Open("postgres", "host=localhost dbname=dealscore sslmode=disable")
	logError(err)
	return db
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
func readJsonConfig() []byte {
	_, filename1, _, _ := runtime.Caller(1)
	fmt.Println(path.Dir(filename1), "$$")
	filename, _ := filepath.Abs(path.Dir(filename1) + "/config.json")
	jsonFile, err := ioutil.ReadFile(filename)
	logError(err)
	data := []byte(jsonFile)
	return data
}
