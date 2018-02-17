package main

import (
    "database/sql"
    "log"
	  _ "github.com/mattn/go-sqlite3"
)

func main() {
  db, err := sql.Open("sqlite3", "./nezha.db")
  if err != nil {
		log.Fatal(err)
	}
  db.Exec(`CREATE TABLE users (
    id INTEGER PRIMARY KEY,
  	name VARCHAR(20),
  	passwd VARCHAR(20),
  	avatar VARCHAR(200),
  	money INTEGER DEFAULT 1000
  ) ;`)
	defer db.Close()
}
