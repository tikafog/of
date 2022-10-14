package boot

import (
	"sync"

	"github.com/cornelk/hashmap"

	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

var (
	once    sync.Once
	modules *hashmap.Map[of.Name, Loader]
)

func init() {
	once.Do(func() {
		modules = hashmap.New[of.Name, Loader]()
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
	if m == nil {
		panic("of: Register module is nil")
	}
	if ok := modules.Insert(m.Name(), m); !ok {
		panic("of: Register called twice for module " + m.Name())
	}
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name) Loader {
	if m, ok := modules.Get(name); ok {
		return m
	}
	return newEmptyLoader(name)
}

// LoadModuleNames ...
// @Description: List all registered modules
// @return []of.Name
func LoadModuleNames() []of.Name {
	names := make([]of.Name, 0, modules.Len())
	modules.Range(func(k of.Name, v Loader) bool {
		names = append(names, of.Name(k))
		return true
	})
	return names
}
