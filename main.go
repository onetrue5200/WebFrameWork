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
		if err = c.WriteJson(http.StatusBadRequest, nil); err != nil {
			panic(err)
		}
	}
	if err = c.WriteJson(http.StatusOK, data); err != nil {
		panic(err)
	}
}

func hello(c *Context) {
	if _, err := fmt.Fprintf(c.W, "hello\n"); err != nil {
		panic(err)
	}
}

func headers(c *Context) {
	for name, headers := range c.R.Header {
		for _, h := range headers {
			if _, err := fmt.Fprintf(c.W, "%v: %v\n", name, h); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	server := NewHttpServer("server", MetricsFilterBuilder)
	server.Route(http.MethodGet, "/hello", hello)
	server.Route(http.MethodGet, "/headers", headers)
	server.Route(http.MethodPost, "/body", body)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
