package of

type CoreModule interface {
	Module(name Name) ModuleStater
}
