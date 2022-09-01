package boot

import (
	"sync"

	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

var (
	initOnce  sync.Once
	modulesMu sync.RWMutex
	modules   map[of.Name]Loader
)

func init() {
	initOnce.Do(func() {
		modules = make(map[of.Name]Loader, 128)
	})
}

type Loader interface {
	of.ModuleStarter
	WithInit(option.InitializeOption) of.ModuleStarter
	WithOption(option.Option) of.ModuleStarter
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

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name) Loader {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m, ok := modules[name]; ok {
		return m
	}
	return newEmptyLoader(name)
}

// LoadModuleNames ...
// @Description: List all registered modules
// @return []of.Name
func LoadModuleNames() []of.Name {
	names := make([]of.Name, 0, len(modules))
	modulesMu.RLock()
	defer modulesMu.RUnlock()
	for name := range modules {
		names = append(names, of.Name(name))
	}
	return names
}
