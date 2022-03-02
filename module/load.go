package module

import (
	"sync"

	"github.com/tikalink/of"
	"github.com/tikalink/of/option"
)

var (
	modulesMu sync.RWMutex
	modules   = initLoadModules()
)

func initLoadModules() [of.NameMax]Loader {
	return [of.NameMax]Loader{
		of.NameAdmin:    newEmptyModule(of.NameAdmin),
		of.NameCenter:   newEmptyModule(of.NameCenter),
		of.NameNode:     newEmptyModule(of.NameNode),
		of.NameBootNode: newEmptyModule(of.NameBootNode),
		of.NameInstruct: newEmptyModule(of.NameInstruct),
	}
}

type Loader interface {
	of.Module
	WithOption(option.InitializeOption) of.Module
}

// Register makes a database module available by the provided name.
// If Register is called twice with the same name or if module is nil,
// it panics.
func Register(m Loader) {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m == nil {
		panic("of: Register module is nil")
	}
	if m.Name() >= of.NameMax {
		panic("of: Register called over module max " + m.Name().String())
	}
	modules[m.Name()] = m
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name, op option.InitializeOption) of.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if name >= of.NameMax {
		panic("of: Load called over module max " + name.String())
	}
	return modules[name].WithOption(op)
}
