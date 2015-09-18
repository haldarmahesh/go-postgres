package database

import (
	"database/sql"
	"encoding/json"
	"github.com/haldarmahesh/dbconnect/helper"
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
	helper.CheckError(err)
	db, err := sql.Open(configuration.Adapter, "host="+configuration.Host+" dbname="+configuration.Dbname+" sslmode="+configuration.Sslmode)
	helper.CheckError(err)
	return db
}

func readJsonConfig() []byte {
	_, filename1, _, _ := runtime.Caller(1)
	filename, _ := filepath.Abs(path.Dir(filename1) + "/config.json")
	jsonFile, err := ioutil.ReadFile(filename)
	helper.CheckError(err)
	data := []byte(jsonFile)
	return data
}
