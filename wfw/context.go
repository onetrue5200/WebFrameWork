package wfw

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer  http.ResponseWriter
	Request *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	if _, err := c.Writer.Write([]byte(fmt.Sprintf(format, value...))); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	if _, err := c.Writer.Write(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if _, err := c.Writer.Write([]byte(html)); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
