package module

type GoWikiPlugin struct {
	routes map[string]func (request HTTPRequest) string
}

func (r GoWikiPlugin) Routes() []string {
	if r.routes == nil {
		return []string{}
	}

	var routes []string
	for k, _ := range r.routes {
		routes = append(routes, k)
	}

	return routes
}

func (r GoWikiPlugin) RegisterRoute(route string, handler func (HTTPRequest) string) {
	if r.routes == nil {
		r.routes = make(map[string]func (request HTTPRequest) string)
	}

	r.routes[route] = handler
}

func (r GoWikiPlugin) HandleRoute(route string, request HTTPRequest) string {
	if r.routes == nil {
		return "Internal server error"
	}

	if routeFunc, ok := r.routes[route]; ok {
		return routeFunc(request)
	}

	return "Not found"
}
