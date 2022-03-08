package of

import (
	"context"
	"encoding/json"

	"github.com/tikafog/of/content"
)

type TypeHandleFunc = func(conn Conn, data json.RawMessage) error

type Core interface {
	Connection
	Event
	Context() context.Context
	State() State
	StoragePath() string

	Inquire(ctx context.Context, r *InquireRequest) error
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
	RegisterEventHandler(from Name, fn TypeEventFunc) error

	//core tools
	Tools() Tools
}

type BootHandler interface {
	AddBoot(ctx context.Context, conn Conn) error
}
