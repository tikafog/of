package dbc

import (
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

func (t *Client[T]) Query() *T {
	return t.c
}

func (t *Client[T]) Update() (*T, *sync.Mutex) {
	t.m.Lock()
	return t.c, &t.m
}

func (t *Client[T]) Close() error {
	return any(t.c).(io.Closer).Close()
}

func openClient[C client](path string, op *Option) (*Client[C], error) {
	var it any = new(C)
	var err error
	var name of.Name
	//var closer io.Closer
	switch it.(type) {
	case *bootnode.Client:
		it, err = openBootNode(of.NameBootnode, path, op)
	case *kernel.Client:
		name = of.NameKernel
		it, err = openKernel(name, path, op)
	case *upgrade.Client:
		name = of.NameUpgrade
		it, err = openUpgrade(name, path, op)
	case *media.Client:
		name = of.NameMedia
		it, err = openMedia(name, path, op)
	default:
		return nil, Error("unsupported client type")
	}
	return &Client[C]{c: (it).(*C)}, err
}
