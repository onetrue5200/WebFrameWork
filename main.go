package main

import (
	"log"
	"net/http"
	"wfw"
)

func main() {
	r := wfw.New()

	r.Use(wfw.Recovery())
	r.GET("/", func(c *wfw.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *wfw.Context) {
		names := []string{""}
		c.String(http.StatusOK, names[100])
	})

	log.Fatal(r.Run(":9999"))
}
