package of

import (
	"bytes"
	"context"
)

type Message struct {
	ID       ID
	Protocol Protocol
	Data     []byte
}

type MessageHandler func(ctx context.Context, conn Conn, buf *bytes.Buffer) error
