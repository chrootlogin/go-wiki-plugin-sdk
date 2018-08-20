package module

import (
	"net/url"
	"encoding/gob"
)

type IGoWikiPlugin interface {
	Name() string
	Version() string
	Routes() []string
	HandleRoute(string, HTTPRequest) string
}

type HTTPRequest struct {
	URL *url.URL
}

type GoWikiPlugin struct {
	Impl IGoWikiPlugin
}

func init() {
	gob.Register(map[string]string{})
	gob.Register(HTTPRequest{})
}