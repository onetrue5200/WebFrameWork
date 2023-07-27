package main

import (
	"log"
	"net/http"
	"wfw"
)

func main() {
	r := wfw.New()

	r.GET("/", func(c *wfw.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *wfw.Context) {
			c.HTML(http.StatusOK, "<h1>Hello V1</h1>")
		})
		v1.GET("/hello", func(c *wfw.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *wfw.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *wfw.Context) {
			c.JSON(http.StatusOK, wfw.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	log.Fatal(r.Run(":9999"))
}
