package of

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/content"
)

type TypeHandleFunc = func(conn Conn, data json.RawMessage) error

type Core interface {
	Connection
	Context() context.Context

	Inquire(ctx context.Context, r *InquireRequest) error
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
	RegisterEventHandler(from Name, fn TypeEventFunc) error
	Event(ctx context.Context, n Name, r *EventRequest) error

	//core tools
	Tools() Tools
}

type BootHandler interface {
	AddBoot(ctx context.Context, conn Conn) error
}
