package of

import (
	"context"
)

type Node interface {
	Connection
	Streamer

	ID() ID
	IsAdmin() bool
	Inquire(ctx context.Context, r *InquireRequest) error
	HandleMessage(protocol Protocol, handler MessageHandler) error
}
