package dbc

import (
	"fmt"

	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/upgrade"
)

type client interface {
	*bootnode.Client | *kernel.Client | *upgrade.Client
}

type OpenFunc[T client] func(name, path string, op *Option) (T, error)

type module[T client] struct {
	name  string
	funcs map[string]OpenFunc[T]
}

func (m module[T]) Name() string {
	return m.name
}

func open[T client](name, path string, op *Option) (T, error) {
	m := module[T]{
		name: name,
		funcs: map[string]OpenFunc[T]{
			"bootnode": openBootNode[T],
			"kernel":   openKernel[T],
			"upgrade":  openUpgrade[T],
		},
	}
	return m.open(path, op)
}

func (m module[T]) open(path string, op *Option) (T, error) {
	v, exist := m.funcs[m.name]
	if exist {
		return v(m.name, path, op)
	}
	return nil, fmt.Errorf("module[%s] not found", m.name)
}
