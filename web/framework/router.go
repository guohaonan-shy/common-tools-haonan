package framework

import (
	"strings"
)

type Router struct {
	router     map[string]HandleFunc
	pathRouter map[string]*TreeNode
}

func NewRouter() *Router {
	return &Router{
		router:     make(map[string]HandleFunc, 0),
		pathRouter: make(map[string]*TreeNode, 0), // key: method
	}
}

func genRouterKey(method string, path string) string {
	builder := strings.Builder{}
	builder.Grow(len(method) + len(path) + len(":"))
	builder.WriteString(method)
	builder.WriteString(path)
	return builder.String()
}

func (r *Router) AddRoute(method, path string, fn HandleFunc) {
	key := genRouterKey(method, path)
	if _, isExist := r.router[key]; isExist {
		panic("same path registered duplicate func")
	}
	//r.router[key] = fn
	// check whether path is existed
	paths := r.parsePath(path)
	if _, ok := r.pathRouter[method]; !ok {
		r.pathRouter[method] = &TreeNode{}
	}

	r.pathRouter[method].insert(path, paths, 0)
	r.router[key] = fn
}

func (r *Router) parsePath(path string) []string {
	pathNodes := strings.Split(path, "/")
	paths := make([]string, 0, len(pathNodes))
	for _, node := range pathNodes {
		if node == "" {
			continue
		}
		paths = append(paths, node)
	}
	return paths
}

//func (r *Router) RouteFunc(method, path string) (HandleFunc, bool) {
//	key := genRouterKey(method, path)
//	fn, isExist := r.router[key]
//	return fn, isExist
//}

func (r *Router) RouteFunc(method string, path string) (HandleFunc, bool) {
	parts := r.parsePath(path)
	root, ok := r.pathRouter[method]

	if !ok {
		return nil, false
	}

	n := root.search(parts, 0)

	if n != nil {
		fn := r.router[genRouterKey(method, path)]
		return fn, true
	}

	return nil, false
}
