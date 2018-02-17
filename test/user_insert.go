package main

import (
    "database/sql"
	  "fmt"
    "log"
    "crypto/md5"
    "encoding/hex"
	  _ "github.com/mattn/go-sqlite3"
)

func main() {
  md5Ctx := md5.New()
  md5Ctx.Write([]byte("123456"))
  cipherStr := md5Ctx.Sum(nil)
  md5_passwd:=hex.EncodeToString(cipherStr)
  db, err := sql.Open("sqlite3", "../nezha.db")
  if err != nil {
		log.Fatal(err)
	}
  stmt, err := db.Prepare("INSERT INTO users(name, passwd, avatar, money)VALUES(?,?, '/avatars/avatar_default.png', 1000);")
  res, err := stmt.Exec("pet", md5_passwd)
  id, err := res.LastInsertId()
	defer db.Close()
	fmt.Println(id)
}
