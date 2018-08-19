package module

import (
	"github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GOWIKI_PLUGIN",
	MagicCookieValue: "gowiki",
}

var PluginMap = map[string]plugin.Plugin{
	"extension": &GoWikiPlugin{},
}