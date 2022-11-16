package of

import (
	"context"
)

type Streamer interface {
	//Receive(ctx context.Context, data *content.Content) error
	Send(ctx context.Context, message *Message) error
}
