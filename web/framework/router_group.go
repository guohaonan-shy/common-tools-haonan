package framework

type RouterGroup struct {
	prefix     string
	middleware []HandleFunc
	//parent     *RouterGroup
	engine *ShyHTTP
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newPrefix := group.prefix + prefix
	newGroup := &RouterGroup{
		prefix: newPrefix,
		//parent: group,
		engine: group.engine,
	}

	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) GetEngine() *ShyHTTP {
	return group.engine
}

func (group *RouterGroup) AddRouteGroup(method, path string, handler HandleFunc) {
	pattern := group.prefix + path
	group.GetEngine().GetRouter().AddRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handleFunc HandleFunc) {
	group.AddRouteGroup("GET", pattern, handleFunc)
}

func (group *RouterGroup) POST(pattern string, handleFunc HandleFunc) {
	group.AddRouteGroup("POST", pattern, handleFunc)
}
