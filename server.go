package main

import "net/http"

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method, pattern string, handleFunc func(c *Context)) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerMapBased()
	root := handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &sdkHttpServer{
		name:    name,
		handler: handler,
		root:    root,
	}
}
