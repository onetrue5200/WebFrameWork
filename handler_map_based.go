package main

import "net/http"

type Routable interface {
	Route(method, pattern string, handleFunc func(c *Context))
}

type Handler interface {
	Routable
	http.Handler
}

type HandlerMapBased struct {
	routers map[string]func(c *Context)
}

func (h *HandlerMapBased) Route(method, pattern string, handleFunc func(c *Context)) {
	key := getKey(method, pattern)
	h.routers[key] = handleFunc
}

func (h *HandlerMapBased) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := getKey(request.Method, request.URL.Path)
	if handler, ok := h.routers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		if _, err := writer.Write([]byte("Page Not Found")); err != nil {
			panic(err)
		}
	}
}

func getKey(method, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerMapBased() Handler {
	return &HandlerMapBased{
		routers: make(map[string]func(c *Context)),
	}
}
