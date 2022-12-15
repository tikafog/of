package of

import (
	"github.com/tikafog/of/feature/bootstrap"
	"github.com/tikafog/of/feature/resource"
)

type CoreFeature interface {
	Module(name Name) Module
	EventTrigger(name Name) EventTrigger
	Node() Node
	Resource() resource.Resource
	Bootstrap() bootstrap.Bootstrap
}
