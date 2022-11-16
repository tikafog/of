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

type AnswerRequest struct {
	Conn   Conn
	Reader io.Reader
}

type Hooker interface {
	Hook(p source.Type, h AdminHookFunc) error
}

type Inquirer interface {
	Inquire(ctx context.Context, r *InquireRequest) error
}

type Answer interface {
	Answer(ctx context.Context, r *AnswerRequest) error
}

type Adminer interface {
	IsAdmin() bool
	Inquirer
	Answer
	Hooker
}
