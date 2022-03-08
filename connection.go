package of

import (
	"context"
)

type Conn interface {
	ID() ID
	LocalAddr() string
	RemoteID() ID
	RemoteAddr() string
}

type Connection interface {
	Send(ctx context.Context, message *Message) error
	Inquire(ctx context.Context, r *InquireRequest) error
	Connect(ctx context.Context, bootaddr ...string) error
	BestPeerConn(id ID) (Conn, error)
}
