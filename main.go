package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/haldarmahesh/dbconnect/database"
	_ "github.com/haldarmahesh/string"
	_ "github.com/lib/pq"
	_ "log"
)

var db *sql.DB

func main() {
	var (
		id   int
		name string
	)
	db = database.CreateConnection()
	rows := readRows()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		logError(err)
		fmt.Println(id, name)
	}
	name, num := readJson()
	fmt.Println(name, num)
	db.Close()
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}
func readRows() *sql.Rows {
	rows, err := db.Query("SELECT id, name from user_details where id=$1", 1)
	logError(err)
	return rows
}
func readJson() (string, int) {
	var data = []byte(`{"status":200, "mahesh" :{"a":12, "b":45}}`)
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("status value: \a", result["mahesh"])
	return "m123", 10
}
