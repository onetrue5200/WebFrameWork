package main

import (
	"fmt"
	"net/http"
)

type Req struct {
	A string `json:"a"`
	B string `json:"b"`
}

func body(w http.ResponseWriter, r *http.Request) {
	c := Context{
		W: w,
		R: r,
	}
	data := &Req{}
	err := c.ReadJson(data)
	if err != nil {
		c.WriteJson(http.StatusBadRequest, nil)
	}
	c.WriteJson(http.StatusOK, data)
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, r *http.Request) {

	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
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
