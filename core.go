package of

import (
	"context"

	"github.com/tikafog/of/buffers/content"
)

type BasicCore interface {
	Context() context.Context
	State() State

	//BasicCore tools
	Tools() Tools
}

type InitCore interface {
	RegisterMessageHandler(protocol Protocol, handler MessageHandler) error
}

type StartCore interface {
	QueryHandler(ct content.Type) (DataQueryHandler, error)
}

// Core
// @Description: the BasicCore of the framework
type Core interface {
	BasicCore
	InitCore
	StartCore
	CoreModule
	Event
	Node() Node
}

//type BootHandler interface {
//	AddBoot(ctx context.Context, conn Conn) error
//}
