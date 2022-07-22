package of

import (
	"context"
)

type Node interface {
	Connection
	ID() ID
	IsAdmin() bool
	Inquire(ctx context.Context, r *InquireRequest) error
}
