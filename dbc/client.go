package dbc

import (
	"errors"

	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type clientAble interface {
	*bootnode.Client | *kernel.Client | *upgrade.Client | *media.Client
	//openBootNode() clientAble
}

type OpenFunc[T clientAble] func(name of.Name, path string, op *Option) (T, error)

type client[T clientAble] struct {
	name of.Name
	path string
	//funcs map[of.Name]OpenFunc[T]
	//openBootNode  OpenFunc[T]
}

func (c client[T]) Name() of.Name { return c.name }

func openClient[T clientAble](path string, op *Option) (T, error) {
	var t T
	var it any = t
	var err error
	var name of.Name
	switch it.(type) {
	case *bootnode.Client:
		name = of.NameBootNode
		it, err = openBootNode[*bootnode.Client](name, path, op)
	case *kernel.Client:
		name = of.NameKernel
		it, err = openKernel[*kernel.Client](name, path, op)
	case *upgrade.Client:
		name = of.NameUpgrade
		it, err = openUpgrade[*upgrade.Client](name, path, op)
	case *media.Client:
		name = of.NameMedia
		it, err = openMedia[*media.Client](name, path, op)
	default:
		return nil, errors.New("unsupported client type")
	}
	return it.(T), err

	//return c.open(op)
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
//		return nil, fmt.Errorf("client[%s] not found", c.name)
//	}
//	return it.(T), err
//	//v, exist := c.funcs[c.name]
//	//if exist {
//	//	return v(c.name, path, op)
//	//}
//
//}
