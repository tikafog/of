package of

import (
	"context"
	"encoding/json"

	"github.com/tikafog/of/content"
)

type TypeHandleFunc = func(conn Conn, data json.RawMessage, args ...Arg) error

type core interface {
	Context() context.Context
	State() State
}

type Core interface {
	core
	CoreModule
	Event

	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error

	Node() Node

	//core tools
	Tools() Tools
}

type BootHandler interface {
	AddBoot(ctx context.Context, conn Conn) error
}
