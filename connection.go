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
	Bootstrap(ctx context.Context, addrs ...string) error
	Connect(ctx context.Context, bootaddr ...string) error
	Disconnect(ctx context.Context, bootaddr ...string) error
	NodeConn(id ID) (Conn, bool)
	NodeLen() int64
	ConnNodes() []Conn
	ConnNodeIDs() []ID
}
