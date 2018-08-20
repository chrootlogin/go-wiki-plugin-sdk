package module

import "net/http"

type GoWikiPlugin struct {
	routes map[string]func (request HTTPRequest) HTTPResponse
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

func (p *GoWikiPlugin) RegisterRoute(route string, handler func (HTTPRequest) HTTPResponse) {
	if p.routes == nil {
		p.routes = make(map[string]func (request HTTPRequest) HTTPResponse)
	}

	p.routes[route] = handler
}

func (p *GoWikiPlugin) HandleRoute(route string, request HTTPRequest) HTTPResponse {
	if p.routes == nil {
		return HTTPResponse{
			Status: http.StatusInternalServerError,
			Body: "Internal server error",
		}
	}

	if routeFunc, ok := p.routes[route]; ok {
		return routeFunc(request)
	}

	return HTTPResponse{
		Status: http.StatusNotFound,
		Body: "Not found",
	}
}
