package of

import (
	"context"
)

type Node interface {
	Connection
	Inquire(ctx context.Context, r *InquireRequest) error
}
