package of

import (
	"context"

	"github.com/tikafog/of/content"
)

type Conn interface {
	ID() ID
	LocalAddr() string
	RemoteID() ID
	RemoteAddr() string
}

type Connection interface {
	Receive(ctx context.Context, data *content.Content) error
	Send(ctx context.Context, message *Message) error
	Connect(ctx context.Context, bootaddr ...string) error
	PeerConn(id ID) (Conn, bool)
	PeerLen() int64
	ConnPeers() []ID
}
