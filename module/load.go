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
	WithInit(option.InitializeOption) of.Module
	WithOption(option.Option) of.Module
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

// Initialize ...
// @Description:
// @param of.Name
// @param option.InitializeOption
// @return of.Module
func Initialize(name of.Name, op option.InitializeOption) of.Module {
	modulesMu.RLock()
	defer modulesMu.RUnlock()
	if name >= of.NameMax {
		panic("of: Initialize called over module max " + name.String())
	}
	return modules[name].WithInit(op)
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name, op option.Option) of.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if name >= of.NameMax {
		panic("of: Load called over module max " + name.String())
	}
	return modules[name].WithOption(op)
}
