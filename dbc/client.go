package dbc

import (
	"fmt"

	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type clientAble interface {
	*bootnode.Client | *kernel.Client | *upgrade.Client | *media.Client
}

type OpenFunc[T clientAble] func(name of.Name, path string, op *Option) (T, error)

type client[T clientAble] struct {
	name  of.Name
	funcs map[of.Name]OpenFunc[T]
}

func (c client[T]) Name() of.Name { return c.name }

//type clientFns[T clientAble] map[string]OpenFunc[T]

func open[T clientAble](name of.Name, path string, op *Option) (T, error) {
	c := client[T]{
		name: name,
		funcs: map[of.Name]OpenFunc[T]{
			of.NameBootNode: openBootNode[T],
			of.NameKernel:   openKernel[T],
			of.NameUpgrade:  openUpgrade[T],
			of.NameMedia:    openMedia[T],
		},
	}
	return c.open(path, op)
}

func (c client[T]) open(path string, op *Option) (T, error) {
	v, exist := c.funcs[c.name]
	if exist {
		return v(c.name, path, op)
	}
	return nil, fmt.Errorf("client[%s] not found", c.name)
}
