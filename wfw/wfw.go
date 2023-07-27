package wfw

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		indexHandler(writer, request)
	case "/hello":
		helloHandler(writer, request)
	default:
		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header [%q] = %q\n", k, v)
	}
}
