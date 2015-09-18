package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "github.com/haldarmahesh/dbconnect/database"
  _ "github.com/haldarmahesh/string"
  _ "github.com/lib/pq"
  "io/ioutil"
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

type As struct {
  As int `json:"as"`
}
type B1 struct {
  A  int `json:"a"`
  B1 As
}
type Mahesh1 struct {
  Mahesh B1
}

func readJson() (string, int) {
  var data = []byte(`{"mahesh" :{"a":12, "b1":{"as":123}}}`)
  /*var result map[string]map[string]interface{}
    if err := json.Unmarshal(data, &result); err != nil {
      fmt.Println("error", err)
    }
    fmt.Println("status value: \a", result["mahesh"]["a"])
  */
  /*var parsed interface{}
    err := json.Unmarshal(data, &parsed)
    logError(err)
    fmt.Println(parsed[:mahesh])
  */
  var mahi Mahesh1
  err := json.Unmarshal(data, &mahi)
  logError(err)
  fmt.Println(mahi.Mahesh.B1)
  readDir()
  return "m123", 10
}
func readDir() {
  files, _ := ioutil.ReadDir("./")
  for _, f := range files {
    fmt.Println(f.Name(), " &&")
  }
}
