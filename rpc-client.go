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

func (p *GoWikiPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &goWikiPluginClient{Broker: b, Client: c}, nil
}

func (p *goWikiPluginClient) Name() string {
	var resp string
	err := p.Client.Call("Plugin.Name", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Name():", resp)
	return resp
}

func (p *goWikiPluginClient) Version() string {
	var resp string
	err := p.Client.Call("Plugin.Version", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Version():", resp)
	return resp
}

func (p *goWikiPluginClient) Routes() []string {
	var resp []string
	err := p.Client.Call("Plugin.Routes", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Version():", resp)
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
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Version():", resp)
	return resp
}