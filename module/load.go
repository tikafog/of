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

func initLoadModules() map[of.Name]Loader {
	return make(map[of.Name]Loader, 128)
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

	if _, ok := modules[m.Name()]; ok {
		panic("of: Register called twice for module " + m.Name())
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
	if m, ok := modules[name]; ok {
		return m.WithInit(op)
	}
	return newEmptyModule(name)
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name, op option.Option) of.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m, ok := modules[name]; ok {
		return m.WithOption(op)
	}
	return newEmptyModule(name)
}
