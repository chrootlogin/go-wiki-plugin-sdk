package module

import (
	"net/rpc"
	"log"

	"github.com/hashicorp/go-plugin"
)

type goWikiPluginClient struct {
	Broker *plugin.MuxBroker
	Client *rpc.Client
}

func (p *GoWikiPluginConnector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &goWikiPluginClient{Broker: b, Client: c}, nil
}

func (p *goWikiPluginClient) Init() {
	err := p.Client.Call("Plugin.Init", new(interface{}), new(interface{}))
	if err != nil {
		log.Fatal(err)
	}
}

func (p *goWikiPluginClient) Name() string {
	var resp string
	err := p.Client.Call("Plugin.Name", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func (p *goWikiPluginClient) Version() string {
	var resp string
	err := p.Client.Call("Plugin.Version", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func (p *goWikiPluginClient) Routes() []string {
	var resp []string
	err := p.Client.Call("Plugin.Routes", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func (p *goWikiPluginClient) HandleRoute(route string, request HTTPRequest) string {
	var resp string
	err := p.Client.Call("Plugin.HandleRoute", map[string]interface{}{
		"route": route,
		"request": request,
	}, &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}