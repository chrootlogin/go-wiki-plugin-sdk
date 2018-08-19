package module

type IGoWikiPlugin interface {
	Name() string
	Version() string
	Routes() []string
	HandleRoute(string) string
}

type GoWikiPlugin struct {
	Impl IGoWikiPlugin
}