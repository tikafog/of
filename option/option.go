package option

import (
	"github.com/tikalink/of"
)

type Option interface {
	Apply(fns ...func(*option) Option)
	ID() of.ID
	Repo() string
	FatherHandler() func(father of.ID)
	Host() interface{}
}

func Repo(repo string) func(o *option) Option {
	return func(o *option) Option {
		o.repo = repo
		return o
	}
}

func FatherHandler(fn func(father of.ID)) func(o *option) Option {
	return func(o *option) Option {
		o.fatherHandler = fn
		return o
	}
}

func Host(host interface{}) func(o *option) Option {
	return func(o *option) Option {
		o.host = host
		return o
	}
}

func ID(id of.ID) func(o *option) Option {
	return func(o *option) Option {
		o.id = id
		return o
	}
}

func Default() Option {
	return &option{}
}

type option struct {
	id            of.ID
	repo          string
	fatherHandler func(father of.ID)
	host          interface{}
}

func (o *option) Apply(fns ...func(*option) Option) {
	for _, fn := range fns {
		fn(o)
	}
}

func (o option) Host() interface{} {
	return o.host
}

func (o option) FatherHandler() func(father of.ID) {
	return o.fatherHandler
}

func (o option) Repo() string {
	return o.repo
}

func (o option) ID() of.ID {
	return o.id
}
