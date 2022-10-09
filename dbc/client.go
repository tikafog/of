package dbc

import (
	"sync"

	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

// ClientType ...
// ENUM(bootnode,kernel,upgrade,media,max)
type ClientType uint32

type client interface {
	kernel.Client | bootnode.Client | upgrade.Client | media.Client
}

type Client[T client] struct {
	Lock sync.Mutex
	C    *T
}

func openClient[C client](path string, op *Option) (*Client[C], error) {
	var it any = new(C)
	var err error
	var name of.Name
	switch it.(type) {
	case *bootnode.Client:
		it, err = openBootNode(of.NameBootNode, path, op)
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
	return &Client[C]{
		C: (it).(*C),
	}, err
}

//c := client[T]{
//	name: name,
//	funcs: map[of.Name]OpenFunc[T]{
//		of.NameBootNode: openBootNode[T],
//		of.NameKernel:   openKernel[T],
//		of.NameUpgrade:  openUpgrade[T],
//		of.NameMedia:    openMedia[T],
//	},
//}

//var fn OpenFunc[T]
//var t T
//switch name {
//case of.NameKernel:
//	fn := openKernel[T]
//case of.NameUpgrade:
//	fn := openUpgrade[T]
//case of.NameBootNode:
//	fn := openBootNode[T]
//case of.NameMedia:
//	fn := openMedia[T]
//}
//openKernel[*kernel.Client](name, path, op)
//return c.open(path, op)
//}
//
//func (c client[T]) open(op *Option) (T, error) {
//	var it any
//	var err error
//	switch c.name {
//	case of.NameBootNode:
//		it, err = openBootNode[*bootnode.Client](c.name, c.path, op)
//	case of.NameKernel:
//		it, err = openKernel[*kernel.Client](c.name, c.path, op)
//	case of.NameUpgrade:
//		it, err = openUpgrade[*upgrade.Client](c.name, c.path, op)
//	case of.NameMedia:
//		it, err = openMedia[*media.Client](c.name, c.path, op)
//	default:
//		return nil, Errorf("client[%s] not found", c.name)
//	}
//	return it.(T), err
//	//v, exist := c.funcs[c.name]
//	//if exist {
//	//	return v(c.name, path, op)
//	//}
//
//}
