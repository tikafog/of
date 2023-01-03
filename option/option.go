package option

import (
	"github.com/tikafog/of"
)

type OptSetFunc func(o *option) Option

type option struct {
	repo  string
	sp    string
	tools of.Tools
}

type Option interface {
	Repo() string
	StoragePath() string
	Tools() of.Tools
}

type ApplyOption interface {
	Option
	Apply(fns ...OptSetFunc)
}

func Tools(tools of.Tools) OptSetFunc {
	return func(o *option) Option {
		o.tools = tools
		return o
	}
}

func Repo(repo string) OptSetFunc {
	return func(o *option) Option {
		o.repo = repo
		return o
	}
}

func StoragePath(sp string) OptSetFunc {
	return func(o *option) Option {
		o.sp = sp
		return o
	}
}

func DefaultApply() ApplyOption {
	return &option{}
}

func (o *option) Apply(fns ...OptSetFunc) {
	for _, fn := range fns {
		fn(o)
	}
}

func (o *option) Repo() string {
	return o.repo
}

func (o *option) StoragePath() string {
	return o.sp
}

func (o *option) Tools() of.Tools {
	return o.tools
}
