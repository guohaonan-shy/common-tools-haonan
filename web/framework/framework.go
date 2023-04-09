package framework

import (
	"net/http"
)

type ShyHTTP struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

type HandleFunc func(ctx *Context)

func NewShyHTTP() *ShyHTTP {
	engine := &ShyHTTP{
		router: NewRouter(),
	}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
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
