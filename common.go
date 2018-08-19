package module

import (
	"github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "GOWIKI_PLUGIN",
	MagicCookieValue: "gowiki",
}

func PluginMap(wikiPlugin IGoWikiPlugin) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"extension": &GoWikiPlugin{
			Impl: wikiPlugin,
		},
	}
}