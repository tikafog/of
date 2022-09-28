package of

import (
	"context"

	"github.com/tikafog/of/content"
)

type Streamer interface {
	Receive(ctx context.Context, data *content.Content) error
	Send(ctx context.Context, message *Message) error
}
