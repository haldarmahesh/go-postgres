package database

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	_ "log"
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
	db, err := sql.Open(configuration.Adapter, "host="+configuration.Host+" dbname="+configuration.Dbname+" sslmode="+configuration.Sslmode)
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
	filename, _ := filepath.Abs(path.Dir(filename1) + "/config.json")
	jsonFile, err := ioutil.ReadFile(filename)
	logError(err)
	data := []byte(jsonFile)
	return data
}
