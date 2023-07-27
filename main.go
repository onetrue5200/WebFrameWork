package main

import (
	"fmt"
	"net/http"
	"wfw"
)

func main() {
	server := wfw.New()
	server.GET("/", indexHandler)
	server.GET("/hello", helloHandler)
	server.Run(":9999")
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header [%q] = %q\n", k, v)
	}
}
