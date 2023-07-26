package main

import (
	"net/http"
	"strings"
)

type HandlerTreeBased struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handleFunc
}

func (h *HandlerTreeBased) Route(method, pattern string, handleFunc handleFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root
	for _, path := range paths {
		child, ok := cur.findChild(path)
		if !ok {
			child = newNode(path)
			cur.children = append(cur.children, child)
		}
		cur = child
	}
	cur.handler = handleFunc
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

func (h *HandlerTreeBased) findRouter(pattern string) (handleFunc, bool) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root
	for _, path := range paths {
		child, ok := cur.findChild(path)
		if !ok {
			return nil, false
		}
		cur = child
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerTreeBased) ServeHTTP(c *Context) {
	handler, ok := h.findRouter(c.R.URL.Path)
	if ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		if _, err := c.W.Write([]byte("Page Not Found")); err != nil {
			panic(err)
		}
	}
}

func NewHandlerTreeBased() Handler {
	return &HandlerTreeBased{
		root: newNode("/"),
	}
}

func (n *node) findChild(path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

var _ Handler = &HandlerTreeBased{}
