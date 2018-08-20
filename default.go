package module

type GoWikiPlugin struct {
	routes map[string]func (request HTTPRequest) string
}

func (r GoWikiPlugin) registerRoute(route string, handler func (HTTPRequest) string) {
	r.routes[route] = handler
}

func (r GoWikiPlugin) HandleRoute(route string, request HTTPRequest) string {

	return "1234"
}
