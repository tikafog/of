package option

import (
	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc"
)

type SetFunc func(o *option) Option

type option struct {
	id   of.ID
	repo string
	sp   string
	//fatherHandler func(father of.ID)
	link  any
	dbc   *dbc.DBC
	tools of.Tools
	//event         of.Event
	//api           of.API
}

type InitializeOption interface {
	Repo() string
	StoragePath() string
	DBC() *dbc.DBC
	Tools() of.Tools
	//Event() of.Event
	//API() of.API
}

type Option interface {
	InitializeOption
	ID() of.ID
	Link() any
	//FatherHandler() func(father of.ID)
}

type ApplyOption interface {
	Option
	Apply(fns ...SetFunc)
}

func Tools(tools of.Tools) SetFunc {
	return func(o *option) Option {
		o.tools = tools
		return o
	}
}

func DBC(dbc *dbc.DBC) SetFunc {
	return func(o *option) Option {
		o.dbc = dbc
		return o
	}
}

//func Event(event of.Event) SetFunc {
//	return func(o *option) Option {
//		o.event = event
//		return o
//	}
//}

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

//func API(api of.API) SetFunc {
//	return func(o *option) Option {
//		o.api = api
//		return o
//	}
//}

//func FatherHandler(fn func(father of.ID)) SetFunc {
//	return func(o *option) Option {
//		o.fatherHandler = fn
//		return o
//	}
//}

func ID(id of.ID) SetFunc {
	return func(o *option) Option {
		o.id = id
		return o
	}
}

func Link(link any) SetFunc {
	return func(o *option) Option {
		o.link = link
		return o
	}
}

func DefaultApply() ApplyOption {
	return &option{}
}

func (o *option) Link() any {
	return o.link
}

func (o *option) Apply(fns ...SetFunc) {
	for _, fn := range fns {
		fn(o)
	}
}

//func (o *option) API() of.API {
//	return o.api
//}
//
//func (o *option) Event() of.Event {
//	return o.event
//}

//func (o *option) FatherHandler() func(father of.ID) {
//	return o.fatherHandler
//}

func (o *option) Repo() string {
	return o.repo
}

func (o *option) StoragePath() string {
	return o.sp
}

func (o *option) ID() of.ID {
	return o.id
}

func (o *option) Tools() of.Tools {
	return o.tools
}

func (o *option) DBC() *dbc.DBC {
	return o.dbc
}
