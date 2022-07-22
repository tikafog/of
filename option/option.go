package option

import (
	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc"
)

type SetFunc func(o *option) Option

type option struct {
	id            of.ID
	repo          string
	sp            string
	fatherHandler func(father of.ID)
	link          interface{}
	dbc           *dbc.DBC
}

func (o *option) DBC() *dbc.DBC {
	return o.dbc
}

type InitializeOption interface {
	Repo() string
	StoragePath() string
	DBC() *dbc.DBC
}

type Option interface {
	InitializeOption
	ID() of.ID
	Link() interface{}
	FatherHandler() func(father of.ID)
}

//type Option interface {
//	StartOption
//}

type ApplyOption interface {
	Option
	Apply(fns ...SetFunc)
}

func DBC(dbc *dbc.DBC) SetFunc {
	return func(o *option) Option {
		o.dbc = dbc
		return o
	}
}

func Repo(repo string) SetFunc {
	return func(o *option) Option {
		o.repo = repo
		return o
	}
}

func StoragePath(sp string) SetFunc {
	return func(o *option) Option {
		o.sp = sp
		return o
	}
}

func FatherHandler(fn func(father of.ID)) SetFunc {
	return func(o *option) Option {
		o.fatherHandler = fn
		return o
	}
}

func ID(id of.ID) SetFunc {
	return func(o *option) Option {
		o.id = id
		return o
	}
}

func Link(link interface{}) SetFunc {
	return func(o *option) Option {
		o.link = link
		return o
	}
}

func DefaultApply() ApplyOption {
	return &option{}
}

func (o *option) Link() interface{} {
	return o.link
}

func (o *option) Apply(fns ...SetFunc) {
	for _, fn := range fns {
		fn(o)
	}
}

func (o option) FatherHandler() func(father of.ID) {
	return o.fatherHandler
}

func (o option) Repo() string {
	return o.repo
}

func (o option) StoragePath() string {
	return o.sp
}

func (o option) ID() of.ID {
	return o.id
}
