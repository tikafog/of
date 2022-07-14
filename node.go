package of

import (
	"context"
)

type Node interface {
	Connection
	ID() ID
	Inquire(ctx context.Context, r *InquireRequest) error
}
