package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/js", "./assets/js")
	r.Static("/css", "./assets/css")
	r.Static("/img", "./assets/img")
	api:= r.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login",
			})
		})
		api.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "login get",
			})
		})
	}

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.tmpl", gin.H{
			"title": "关于",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "哪吒",
		})
	})

	r.Run(":3000")// listen and serve on 0.0.0.0:8080
}
