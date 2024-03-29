package of

import (
	"encoding/json"
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
	NameSystem      = OwnName(7852121090758537921, "system")
	NameAdmin       = OwnName(2239877576446804528, "admin")
	NameAdminClient = OwnName(14909113260478095105, "admin_client")
	NameCenter      = OwnName(18108201700337443927, "center")
	NameNode        = OwnName(14580695937169233336, "node")
	NameBootnode    = OwnName(16813717775905332386, "bootnode")
	NameUpgrade     = OwnName(10629397049073578029, "upgrade")
	NameMedia       = OwnName(2061340576534510319, "media")
	NameKernel      = OwnName(15590474080070586392, "kernel")
	NameInstruct    = OwnName(3635930719169311595, "instruct")
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

type ModuleConfigLoader interface {
	LoadConfig(json.RawMessage) (json.RawMessage, error)
}

//type ModuleStarter interface {
//	ModuleConfiger
//ModuleRunner
//
//	//this all calls before run
//	CoreRegister
//
//	//this all calls after run
//	//APIRegister
//	//EventRegister
//
//Module
//}

type ModuleInjector interface {
	Inject(name string, v interface{}) error
}

type Module interface {
	Name() Name
	IsValid() bool
}

// CompareName compares the name of a module
// @param Name
// @param Name
// @return bool
func CompareName(n, o Name) bool {
	return n.ID() == o.ID()
}

// CompareNameNeq compares the name when the name is eq other one returns true
// @param Name
// @param ...Name
// @return bool
func CompareNameNeq(n Name, others ...Name) bool {
	for i := range others {
		if n.ID() != others[i].ID() {
			return false
		}
	}
	return true
}

// CompareNameEq compares the name when the name is neq other one returns true
// @param Name
// @param ...Name
// @return bool
func CompareNameEq(n Name, others ...Name) bool {
	for i := range others {
		if n.ID() == others[i].ID() {
			return true
		}
	}
	return false
}

// OwnName create self name with id
// @param uint64
// @param string
// @return Name
func OwnName(id uint64, value string) Name {
	return &name{
		id:    id,
		value: value,
	}
}
