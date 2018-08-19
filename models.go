package module

// IGreeter is the interface that is known between the mainapp and plugins
type IGoWikiPlugin interface {
	Name() string
	Version() string
	Routes() []string
}

// Greeter is a structure that has a reference to the interface
type GoWikiPlugin struct {
	Impl IGoWikiPlugin
}