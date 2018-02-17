package main

import (
    "database/sql"
	  "fmt"
    "log"
	  _ "github.com/mattn/go-sqlite3"
)

func main() {
  db, err := sql.Open("sqlite3", "../nezha.db")
  if err != nil {
		log.Fatal(err)
	}
  rows, err := db.Query("select id,name from users")
  for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
  }
	defer db.Close()
}
