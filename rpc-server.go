package module

import (
	"github.com/hashicorp/go-plugin"
	"errors"
)

type goWikiPluginServer struct {
	Broker   	  *plugin.MuxBroker
	IGoWikiPlugin IGoWikiPlugin
}

// Server implmentation of go-plugin.plugin.Plugin.Server
func (p *GoWikiPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	if p.Impl == nil {
		return nil, errors.New("GoWikiPlugin interface not implemeted")
	}
	return &goWikiPluginServer{Broker: b, IGoWikiPlugin: p.Impl()}, nil
}

////////////////////////////////////////////////////////////////////////////////
//
//  For each function defined in the interface, there needs to be an equivalent
//  greeterServer call
//
////////////////////////////////////////////////////////////////////////////////

// Name calls the plugin implementation to get the name of the plugin
func (p *goWikiPluginServer) Name(nothing interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Name()
	return nil
}

// Version calls the plugin implementation to get the version of the plugin
func (p *goWikiPluginServer) Version(nothing interface{}, result *string) error {
	*result = p.IGoWikiPlugin.Version()
	return nil
}