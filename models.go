package module

import "net/url"

type IGoWikiPlugin interface {
	Name() string
	Version() string
	Routes() []string
	HandleRoute(string, HTTPRequest) string
}

type HTTPRequest struct {
	URL url.URL
}

type GoWikiPlugin struct {
	Impl IGoWikiPlugin
}