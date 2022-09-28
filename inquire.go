package of

import (
	"context"
)

type InquireRequest struct {
	Data []byte
}

type Inquirer interface {
	IsAdmin() bool
	Inquire(ctx context.Context, r *InquireRequest) error
}
