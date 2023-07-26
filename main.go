package main

import (
	"fmt"
	"net/http"
)

type Req struct {
	A string `json:"a"`
	B string `json:"b"`
}

func body(c *Context) {
	data := &Req{}
	err := c.ReadJson(data)
	if err != nil {
		c.WriteJson(http.StatusBadRequest, nil)
	}
	c.WriteJson(http.StatusOK, data)
}

func hello(c *Context) {
	fmt.Fprintf(c.W, "hello\n")
}

func headers(c *Context) {
	for name, headers := range c.R.Header {
		for _, h := range headers {
			fmt.Fprintf(c.W, "%v: %v\n", name, h)
		}
	}
}

func main() {
	server := NewHttpServer("server")
	server.Route("/hello", hello)
	server.Route("/headers", headers)
	server.Route("/body", body)
	server.Start(":8080")
}
