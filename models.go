package module

import (
	"net/url"
	"encoding/gob"
)

type IGoWikiPlugin interface {
	Init()
	Name() string
	Version() string
	Routes() []string
	HandleRoute (string, HTTPRequest) HTTPResponse
}

type HTTPRequest struct {
	URL *url.URL
}

type HTTPResponse struct {
	Status int
	Headers map[string]string
	Body string
}

type GoWikiPluginConnector struct {
	Impl IGoWikiPlugin
}

func init() {
	gob.RegisterName("HTTPRequest", HTTPRequest{})
	gob.RegisterName("HTTPResponse", HTTPResponse{})
}