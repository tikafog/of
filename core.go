package of

import (
	"context"

	"github.com/tikafog/of/content"
)

type core interface {
	Context() context.Context
	State() State
}

type Core interface {
	core
	CoreModule
	Event

	RegisterDataHandler(ct content.Type, handler DataHandler) error

	Node() Node

	//core tools
	Tools() Tools
}

//type BootHandler interface {
//	AddBoot(ctx context.Context, conn Conn) error
//}
