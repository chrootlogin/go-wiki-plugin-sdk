package module

type GoWikiPlugin struct {
	routes map[string]func (request HTTPRequest) string
}

func (p *GoWikiPlugin) Routes() []string {
	if p.routes == nil {
		return []string{}
	}

	var routes []string
	for k, _ := range p.routes {
		routes = append(routes, k)
	}

	return routes
}

func (p *GoWikiPlugin) RegisterRoute(route string, handler func (HTTPRequest) string) {
	if p.routes == nil {
		p.routes = make(map[string]func (request HTTPRequest) string)
	}

	p.routes[route] = handler
}

func (p *GoWikiPlugin) HandleRoute(route string, request HTTPRequest) string {
	if p.routes == nil {
		return "Internal server error"
	}

	if routeFunc, ok := p.routes[route]; ok {
		return routeFunc(request)
	}

	return "Not found"
}
