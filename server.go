package main

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r)
		handlerFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
