package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "database/sql"
	  //"fmt"
	  _ "github.com/mattn/go-sqlite3"
	  "log"
	  //"os"
)

func main() {
	r:= gin.Default()
  db, err := sql.Open("sqlite3", "./nezha.db")
  if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./assets/js")
	r.Static("/css", "./assets/css")
	r.Static("/img", "./assets/img")
	api:= r.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
      query_name := c.PostForm("name")
      query_passwd := c.PostForm("passwd")

      msg:=""
      //
      sql:="select id,name,passwd from users where name='"+query_name+"'"

      rows, err := db.Query(sql)

      if rows.Next() {
    		var id int
    		var name string
        var passwd string
    		err = rows.Scan(&id, &name,&passwd)
    		if err != nil {
    			log.Fatal(err)
    		}
        if(passwd==query_passwd){
          msg="登录成功"
        }else{
          msg="登录失败"
        }
      }else{
        msg="账户不存在"
      }
      //

			c.JSON(200, gin.H{
				"message": msg,
			})
		})
		api.POST("/register", func(c *gin.Context) {
      query_name := c.PostForm("name")
      query_passwd := c.PostForm("passwd")

      msg:=""
      //
      sql:="select id,name,passwd from users where name='"+query_name+"'"

      rows, err := db.Query(sql)

      if rows.Next() {
        msg="用户已经存在"
        c.JSON(200, gin.H{
          "err": err,
          "message": msg,
        })
      }else{
        stmt, err := db.Prepare("INSERT INTO users(name, passwd, avatar, money)VALUES(?,?, '/avatars/avatar_default.png', 1000);")
        res, err := stmt.Exec(query_name, query_passwd)
        id, err := res.LastInsertId()
        msg="注册成功"
        c.JSON(200, gin.H{
          "err": err,
          "message": msg,
          "id": id,
        })
      }
      //


		})
	}

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tmpl", gin.H{
			"title": "关于",
		})
	})

  r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title": "登录",
		})
	})

  r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title": "注册",
		})
	})

  r.GET("/logout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "logout.tmpl", gin.H{
			"title": "您已经安全退出了",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "哪吒",
		})
	})

  r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "哪吒区块链",
		})
	})

  r.GET("/users", func(c *gin.Context) {

		c.HTML(http.StatusOK, "users.tmpl", gin.H{
			"title": "用户列表",
		})
	})



	r.Run(":80")// listen and serve on 0.0.0.0:8080
}
