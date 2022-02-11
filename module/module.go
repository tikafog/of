package module

import (
	"github.com/tikalink/of"
)

//Name returns the names of all the modules
//ENUM(bootstrap,instruct,max)
type Name int

type TypeHandleFunc = func(id of.ID, data []byte) error

type Module interface {
	Start() error
	Name() Name
	TypeHandleFunc
}
