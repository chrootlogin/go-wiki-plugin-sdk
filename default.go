package module

type GoWikiPlugin struct {
	routes map[string]func (request HTTPRequest) string
}

func (r GoWikiPlugin) RegisterRoute(route string, handler func (HTTPRequest) string) {
	if r.routes == nil {
		r.routes = make(map[string]func (request HTTPRequest) string)
	}
	r.routes[route] = handler
}

func (r GoWikiPlugin) HandleRoute(route string, request HTTPRequest) string {

	return "1234"
}
