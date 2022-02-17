package of

import (
	"context"
)

type Core interface {
	Context() context.Context
	Send(ctx context.Context, message *Message) error
}
