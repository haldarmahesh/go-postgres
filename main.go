package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	var (
		id   int
		name string
	)
	db, err := sql.Open("postgres", "host=localhost dbname=dealscore sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conected, time is %s", time.Now())
	rows, err := db.Query("SELECT id, name from user_details where id =$1", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	db.Close()
}
