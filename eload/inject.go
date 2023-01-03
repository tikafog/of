package eload

import (
	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc"
)

type Injector interface {
	Inject(v any) error
}

type DatabaseRegister interface {
	RegisterDBC(dbc *dbc.DBC) error
}

type EventRegister interface {
	RegisterEvent(of.Event) error
}

type APIRegister interface {
	RegisterAPI(of.API) error
}
