package module

// IGreeter is the interface that is known between the mainapp and plugins
type IGoWikiPlugin interface {
	Name() string
	Version() string
	StartTime() int64
}

// Greeter is a structure that has a reference to the interface
type GoWikiPlugin struct {
	Impl IGoWikiPlugin
}