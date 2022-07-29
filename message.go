package of

import (
	"context"
	"io"
)

type Message struct {
	ID       ID
	Protocol Protocol
	Data     []byte
}

type MessageHandler func(ctx context.Context, conn Conn, reader io.Reader) error
