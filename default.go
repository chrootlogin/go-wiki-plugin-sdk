package module

type GoWikiPlugin struct {}

func (r GoWikiPlugin) HandleRoute(route string, request HTTPRequest) string {
	return "1234"
}
