package of

import (
	"context"
)

type Node interface {
	Connection
	ID() ID
	Impl(module Module) error
	IsAdmin() bool
	Inquire(ctx context.Context, r *InquireRequest) error
}
