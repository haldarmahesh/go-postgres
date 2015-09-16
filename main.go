package main

import (
  "fmt"
  _ "github.com/haldarmahesh/string"
  _ "github.com/lib/pq"
  _ "log"
)

import "github.com/haldarmahesh/dbconnect/database"
func main() {
  var (
    id   int
    name string
  )
  db := database.CreateConnection()
  rows, err := db.Query("SELECT id, name from user_details where id=$1", 1)
  logError(err)
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&id, &name)
    logError(err)
    fmt.Println(id, name)
  }
  db.Close()

  
}

func logError(err error) {
  if err != nil {
    panic(err)
  }
}
