package of

import (
	"context"

	"github.com/tikafog/of/content"
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

	RegisterDataHandler(ct content.Type, handler DataHandler[any, any]) error

	Node() Node

	//core tools
	Tools() Tools
}

//type BootHandler interface {
//	AddBoot(ctx context.Context, conn Conn) error
//}
