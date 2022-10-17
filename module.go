package of

import (
	"context"
)

// Name returns the names of all the modules
type Name interface {
	Compare(other Name) bool
	String() string
	ID() uint64
}

type name struct {
	id    uint64
	value string
}

var (
	NameNotSet      = OwnName(0, "notset")
	NameAccount     = OwnName(18275347644319189965, "account")
	NameAdmin       = OwnName(2239877576446804528, "admin")
	NameAdminClient = OwnName(14909113260478095105, "admin_client")
	NameCenter      = OwnName(18108201700337443927, "center")
	NameNode        = OwnName(14580695937169233336, "node")
	NameBootnode    = OwnName(16813717775905332386, "bootnode")
	NameUpgrade     = OwnName(10629397049073578029, "upgrade")
	NameMedia       = OwnName(2061340576534510319, "media")
	NameKernel      = OwnName(15590474080070586392, "kernel")
	NameGateway     = OwnName(6483280667709159949, "gateway")
)

func (n name) Compare(other Name) bool {
	return n.ID() == other.ID()
}

func (n name) String() string {
	return n.value
}

func (n name) ID() uint64 {
	return n.id
}

type ModuleStarter interface {
	Init() error
	Run(ctx context.Context) error
	Destroy()

	//this all calls before run
	PreloadCore(core Core) error

	//this all calls after run
	RegisterAPI(api API) error
	RegisterEvent(event Event) error
	//RegisterMessageHandler(name string, handler MessageHandler) error
	//this all calls after init
	//RegisterMessageHandler(name string, handler MessageHandler) error

	Module
}

type Module interface {
	Name() Name
	IsValid() bool
}

func CompareName(n, o Name) bool {
	return n.String() == o.String()
}

func CompareNameNeq(n name, others ...Name) bool {
	for i := range others {
		if n.ID() != others[i].ID() {
			return false
		}
	}
	return true
}

func CompareNameEq(n name, others ...Name) bool {
	for i := range others {
		if n.ID() == others[i].ID() {
			return true
		}
	}
	return false
}

func OwnName(id uint64, value string) Name {
	return &name{
		id:    id,
		value: value,
	}
}
