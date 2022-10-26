package of

import (
	"context"
	"io"
)

type InquireRequest struct {
	Data []byte
}

type AnswerRequest struct {
	conn   Conn
	reader io.Reader
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
}
