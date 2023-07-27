package main

import (
	"log"
	"net/http"
	"wfw"
)

func main() {
	server := wfw.New()
	server.GET("/", indexHandler)
	server.GET("/hello", helloHandler)
	server.POST("/login", loginHandler)
	log.Fatal(server.Run(":9999"))
}

func indexHandler(c *wfw.Context) {
	c.HTML(http.StatusOK, "<h1>Hello World</h1>")
}

func helloHandler(c *wfw.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *wfw.Context) {
	c.JSON(http.StatusOK, wfw.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}
