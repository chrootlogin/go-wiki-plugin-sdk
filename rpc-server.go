package module

import (
	"github.com/hashicorp/go-plugin"
	"errors"
	"net/url"
	"log"
)

type goWikiPluginServer struct {
	Broker   	  *plugin.MuxBroker
	IGoWikiPlugin IGoWikiPlugin
}

func (p *GoWikiPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	if p.Impl == nil {
		return nil, errors.New("GoWikiPlugin interface not implemeted")
	}
	return &goWikiPluginServer{Broker: b, IGoWikiPlugin: p.Impl}, nil
}

func (p *goWikiPluginServer) Name(nothing interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Name()
	return nil
}

func (p *goWikiPluginServer) Version(nothing interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Version()
	return nil
}

func (p *goWikiPluginServer) Routes(nothing interface{}, result *[]string) error {
	*result = p.IGoWikiPlugin.Routes()
	return nil
}

func (p *goWikiPluginServer) HandleRoute(args map[string]interface{}, result *string) error {
	route := args["route"].(string)
	request := args["request"].(map[string]string)

	reqUrl, err := url.Parse(request["url"])
	if err != nil {
		log.Fatal(err)
	}

	httpRequest := HTTPRequest{
		URL: reqUrl,
	}

	*result = p.IGoWikiPlugin.HandleRoute(route, httpRequest)
	return nil
}