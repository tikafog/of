package boot

import (
	"sync"

	"github.com/cornelk/hashmap"

	"github.com/tikafog/of"
	"github.com/tikafog/of/option"
)

var (
	once    sync.Once
	names   *hashmap.Map[uint64, of.Name]
	modules *hashmap.Map[uint64, Loader]
)

func init() {
	once.Do(func() {
		modules = hashmap.New[uint64, Loader]()
		names = hashmap.New[uint64, of.Name]()
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
	if ok := names.Insert(m.Name().ID(), m.Name()); !ok {
		panic("of: Register called twice for module " + m.Name().String())
	}
	modules.Set(m.Name().ID(), m)
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name) Loader {
	if m, ok := modules.Get(name.ID()); ok {
		return m
	}
	return newEmptyLoader(name)
}

// LoadModuleNames ...
// @Description: List all registered modules
// @return []of.Name
func LoadModuleNames() []of.Name {
	rets := make([]of.Name, 0, modules.Len())
	modules.Range(func(k uint64, v Loader) bool {
		getted, ok := names.Get(k)
		if ok {
			rets = append(rets, getted)
		}
		return true
	})
	return rets
}
