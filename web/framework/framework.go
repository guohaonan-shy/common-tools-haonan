package framework

import (
	"fmt"
	"net/http"
)

type ShyHTTP struct {
	router map[string]*Entry
}

func NewShyHTTP() *ShyHTTP {
	return &ShyHTTP{
		router: make(map[string]*Entry, 0),
	}
}

type HTTPFunc func(resp http.ResponseWriter, req *http.Request)

type Entry struct {
	fn     HTTPFunc
	method string
}

func NewEntry(method string, fn HTTPFunc) *Entry {
	return &Entry{
		fn:     fn,
		method: method,
	}
}

func (entry *Entry) Match(method string) bool {
	return method == entry.method
}

func (engine *ShyHTTP) GET(path string, httpFunc HTTPFunc) {
	entry := NewEntry("GET", httpFunc)
	if _, ok := engine.router[path]; ok {
		panic("same path registered duplicate func")
	}

	engine.router[path] = entry
}

func (engine *ShyHTTP) POST(path string, httpFunc HTTPFunc) {
	entry := NewEntry("POST", httpFunc)
	if _, ok := engine.router[path]; ok {
		panic("same path registered duplicate func")
	}

	engine.router[path] = entry
}

func (engine *ShyHTTP) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	if entry, isMatch := engine.router[req.URL.Path]; isMatch {
		if entry.Match(req.Method) {
			entry.fn(resp, req)
		} else {
			fmt.Fprintf(resp, "404: method hasn't matched:%s\n", req.Method)
		}
	} else {
		fmt.Fprintf(resp, "404: path hasn't registered:%s\n", req.URL)
	}
}

func (engine *ShyHTTP) Run(addr string) {
	http.ListenAndServe(addr, engine)
}
