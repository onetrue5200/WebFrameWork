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

	r.GET("/hello", func(c *wfw.Context) {
		c.String(http.StatusOK, "hello %s, your're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *wfw.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *wfw.Context) {
		c.JSON(http.StatusOK, wfw.H{"filepath": c.Param("filepath")})
	})

	log.Fatal(r.Run(":9999"))
}
