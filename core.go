package of

import (
	"context"
)

type Core interface {
	Context() context.Context
	SendRequest(ctx context.Context, ID, data []byte) error
}
