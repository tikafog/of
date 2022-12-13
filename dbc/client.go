package dbc

import (
	"context"
	"io"
	"sync"

	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type client interface {
	kernel.Client | bootnode.Client | upgrade.Client | media.Client
}

type Client[T client] struct {
	m sync.Mutex
	c *T
}

// DB returns the database connection
// @receiver *Client[T]
// @return *T
func (t *Client[T]) DB() *T {
	return t.c
}

// UpdateStart lock for database updates
// @receiver *Client[T]
// @param context.Context
// @return context.Context
func (t *Client[T]) UpdateStart(ctx context.Context) context.Context {
	return context.WithValue(ctx, "privacy", &t.m)
}

// UpdateFinish unlock for database updates
// @receiver *Client[T]
func (t *Client[T]) UpdateFinish() {
	t.m.Unlock()
}

func (t *Client[T]) Close() error {
	return any(t.c).(io.Closer).Close()
}

func openClient[C client](path string, op *Option) (*Client[C], error) {
	var it any = new(C)
	var err error
	//var name of.Name
	//var closer io.Closer
	switch it.(type) {
	case *bootnode.Client:
		//name = of.NameBootnode
		it, err = openBootNode(of.NameBootnode, path, op)
	case *kernel.Client:
		//name = of.NameKernel
		it, err = openKernel(of.NameKernel, path, op)
	case *upgrade.Client:
		//name = of.NameUpgrade
		it, err = openUpgrade(of.NameUpgrade, path, op)
	case *media.Client:
		//name = of.NameMedia
		it, err = openMedia(of.NameMedia, path, op)
	default:
		return nil, Error("unsupported client type")
	}
	return &Client[C]{c: (it).(*C)}, err
}
