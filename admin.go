package of

import (
	"context"
	"io"

	"github.com/tikafog/of/feature/source"
)

type AdminHookFunc = func(ctx context.Context, data []byte) error

type InquireRequest struct {
	Data []byte
}

type AdminDataRequest struct {
	Conn   Conn
	Reader io.Reader
}

type Hooker interface {
	Hook(p source.Type, h AdminHookFunc) error
}

type Inquirer interface {
	Inquire(ctx context.Context, r *InquireRequest) error
}

type StreamCaller interface {
	SetStreams(streams [ProtocolMax]MessageHandler)
	Stream(protocol Protocol) (MessageHandler, bool)
}

type Adminer interface {
	IsAdmin() bool
	Inquirer
	StreamCaller
	Hooker
}
