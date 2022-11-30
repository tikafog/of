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

type AdminHandler interface {
	Receive(ctx context.Context, protocol Protocol, r *AdminDataRequest) error
	SetStreams(streams [ProtocolMax]MessageHandler)
}

type Adminer interface {
	IsAdmin() bool
	Inquirer
	AdminHandler
	Hooker
}
