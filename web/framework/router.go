package framework

import (
	"strings"
)

type HandleFunc func(ctx *Context)
type Router struct {
	router map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		router: make(map[string]HandleFunc, 0),
	}
}

func genRouterKey(method string, path string) string {
	builder := strings.Builder{}
	builder.Grow(len(method) + len(path) + len(":"))
	builder.WriteString(method)
	builder.WriteString(path)
	builder.WriteString(":")
	return builder.String()
}

func (r *Router) AddRoute(method, path string, fn HandleFunc) {
	key := genRouterKey(method, path)
	if _, isExist := r.router[key]; isExist {
		panic("same path registered duplicate func")
	}
	r.router[key] = fn
}

func (r *Router) RouteFunc(method, path string) (HandleFunc, bool) {
	key := genRouterKey(method, path)
	fn, isExist := r.router[key]
	return fn, isExist
}
