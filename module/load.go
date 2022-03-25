package module

import (
	"sync"

	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

var (
	modulesMu sync.RWMutex
	modules   = initLoadModules()
)

func initLoadModules() map[string]Loader {
	return make(map[string]Loader, 128)
}

type Loader interface {
	of.ModuleStarter
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

	if _, ok := modules[m.Name().String()]; ok {
		panic("of: Register called twice for module " + m.Name())
	}
	modules[m.Name().String()] = m
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name) Loader {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m, ok := modules[name.String()]; ok {
		return m
	}
	return newEmptyLoader(name)
}

// ListLoaders ...
// @Description: List all registered modules
// @return []of.Name
func ListLoaders() []of.Name {
	names := make([]of.Name, 0, len(modules))
	modulesMu.RLock()
	defer modulesMu.RUnlock()
	for name := range modules {
		names = append(names, of.Name(name))
	}
	return names
}
