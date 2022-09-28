package of

type CoreFeature interface {
	Module(name Name) Module
	API() API
	Event() Event
	Node() Node
	Resource() Resource
	Bootstrap() Bootstrap
}
