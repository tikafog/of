package module

import (
	"encoding/json"

	"github.com/tikalink/of/module/option"
)

//Name returns the names of all the modules
//ENUM(bootNode,instruct,max)
type Name int

type TypeHandleFunc = func(id string, data json.RawMessage) error

type Module interface {
	Valid() bool
	Start() error
	Name() Name
	Option(op option.Option) error
	HandleData(id string, data json.RawMessage) error
}
