package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var (
		id   int
		name string
	)
	db := createDbConnection()
	rows, err := db.Query("SELECT id, name from user_details where id =$1", 1)
	logError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		logError(err)
		fmt.Println(id, name)
	}
	db.Close()
}

func createDbConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost dbname=dealscore sslmode=disable")
	logError(err)
	return db
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
