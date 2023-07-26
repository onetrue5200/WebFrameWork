package main

import "net/http"

type Routable interface {
	Route(method, pattern string, handleFunc func(c *Context))
}

type Handler interface {
	Routable
	ServeHTTP(c *Context)
}

type HandlerMapBased struct {
	routers map[string]func(c *Context)
}

func (h *HandlerMapBased) Route(method, pattern string, handleFunc func(c *Context)) {
	key := h.getKey(method, pattern)
	h.routers[key] = handleFunc
}

func (h *HandlerMapBased) ServeHTTP(c *Context) {
	key := h.getKey(c.R.Method, c.R.URL.Path)
	if handler, ok := h.routers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		if _, err := c.W.Write([]byte("Page Not Found")); err != nil {
			panic(err)
		}
	}
}

func (h *HandlerMapBased) getKey(method, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerMapBased() Handler {
	return &HandlerMapBased{
		routers: make(map[string]func(c *Context)),
	}
}
