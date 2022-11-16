package of

import (
	"context"
	"io"

	"github.com/tikafog/of/feature/source"
)

type AdminHooker = func(ctx context.Context, data []byte) error

type InquireRequest struct {
	Data []byte
}

type AnswerRequest struct {
	Conn   Conn
	Reader io.Reader
}

type Inquirer interface {
	Inquire(ctx context.Context, r *InquireRequest) error
	Hook(p source.Type, h AdminHooker) error
}

type Answer interface {
	Answer(ctx context.Context, r *AnswerRequest) error
}

type Adminer interface {
	IsAdmin() bool
	Inquirer
	Answer
}
