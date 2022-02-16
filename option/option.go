package option

import (
	"github.com/tikalink/of"
)

type Option interface {
	Repo() string
	FatherHandler() func(father of.ID)
	Host() interface{}
}

func Repo(repo string) func(o Option) Option {
	return func(o Option) Option {
		op := o.(option)
		op.repo = repo
		return op
	}
}

func FatherHandler(fn func(father of.ID)) func(o Option) Option {
	return func(o Option) Option {
		op := o.(option)
		op.fatherHandler = fn
		return op
	}
}

func Host(host interface{}) func(o Option) Option {
	return func(o Option) Option {
		op := o.(option)
		op.host = host
		return op
	}
}

func Default() Option {
	return &option{}
}

type option struct {
	repo          string
	fatherHandler func(father of.ID)
	host          interface{}
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
