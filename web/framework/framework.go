package framework

import (
	"net/http"
)

type ShyHTTP struct {
	router *Router
}

type HandleFunc func(ctx *Context)

func NewShyHTTP() *ShyHTTP {
	return &ShyHTTP{
		router: NewRouter(),
	}
}

func (engine *ShyHTTP) GET(path string, fn HandleFunc) {
	engine.router.AddRoute("GET", path, fn)
}

func (engine *ShyHTTP) POST(path string, fn HandleFunc) {
	engine.router.AddRoute("POST", path, fn)
}

func (engine *ShyHTTP) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	if fn, isMatch := engine.router.RouteFunc(req.Method, req.URL.Path); isMatch {
		fn(NewContext(resp, req))
	} else {
		http.Error(resp, "404: path hasn't registered\n", 404)
	}
}

func (engine *ShyHTTP) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
