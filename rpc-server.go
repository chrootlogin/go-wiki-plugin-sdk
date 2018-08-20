package module

import (
	"github.com/hashicorp/go-plugin"
	"errors"
		)

type goWikiPluginServer struct {
	Broker   	  *plugin.MuxBroker
	IGoWikiPlugin IGoWikiPlugin
}

func (p *GoWikiPluginConnector) Server(b *plugin.MuxBroker) (interface{}, error) {
	if p.Impl == nil {
		return nil, errors.New("GoWikiPlugin interface not implemeted")
	}
	return &goWikiPluginServer{Broker: b, IGoWikiPlugin: p.Impl}, nil
}

func (p *goWikiPluginServer) Init(args interface{}, result interface{}) error {
	p.IGoWikiPlugin.Init()

	return nil
}

func (p *goWikiPluginServer) Name(args interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Name()
	return nil
}

func (p *goWikiPluginServer) Version(args interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Version()
	return nil
}

func (p *goWikiPluginServer) Routes(args interface{}, result *[]string) error {
	*result = p.IGoWikiPlugin.Routes()
	return nil
}

func (p *goWikiPluginServer) HandleRoute(args map[string]interface{}, result *string) error {
	route := args["route"].(string)
	request := args["request"].(HTTPRequest)

	*result = p.IGoWikiPlugin.HandleRoute(route, request)
	return nil
}