package wfw

import (
	"html/template"
	"net/http"
	"strings"
)

type Engine struct {
	*RouteGroup
	router        *router
	groups        []*RouteGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(request.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(writer, request)
	c.handlers = middlewares
	c.enginer = e
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
