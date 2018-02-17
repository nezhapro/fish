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
  sql:="select id,name,passwd from users where name='tom'"

  rows, err := db.Query(sql)

  if rows.Next() {
		var id int
		var name string
    var passwd string
		err = rows.Scan(&id, &name,&passwd)
		if err != nil {
			log.Fatal(err)
		}
    if(passwd=="e10adc3949ba59abbe56e057f20f883e"){
      fmt.Println("登录成功")
    }else{
      fmt.Println("登录失败")
    }
  }else{
    fmt.Println("账户不存在")
  }
	defer db.Close()
}
