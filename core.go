package of

import (
	"context"

	"github.com/tikafog/of/buffers/content"
)

type core interface {
	Context() context.Context
	State() State
}

// Core
// @Description: the core of the framework
type Core interface {
	core
	CoreModule
	Event

	QueryHandler(ct content.Type) (DataQueryHandler, error)

	Node() Node

	//core tools
	Tools() Tools
}

//type BootHandler interface {
//	AddBoot(ctx context.Context, conn Conn) error
//}
