package main

import (
	"log"
	"net/http"
	"time"
	"wfw"
)

func onlyForV2() wfw.HandlerFunc {
	return func(c *wfw.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := wfw.New()

	r.Use(wfw.Logger())
	r.GET("/", func(c *wfw.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *wfw.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	log.Fatal(r.Run(":9999"))
}
