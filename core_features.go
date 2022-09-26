package of

type CoreFeature interface {
	Module(name Name) Module
	Event() Event
	Node() Node
	Resource() Resource
}
