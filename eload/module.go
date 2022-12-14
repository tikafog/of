package eload

import (
	"github.com/tikafog/of"
	"sync"
)

var (
	_once   sync.Once
	_loader Loader
	//names   *hashmap.Map[uint64, of.Name]
	//modules *hashmap.Map[uint64, of.ModuleInterface]
)

func init() {
	_once.Do(func() {
		_loader = newLoader()
		//modules = hashmap.New[uint64, of.ModuleInterface]()
		//names = hashmap.New[uint64, of.Name]()
	})
}

// Register makes a database module available by the provided name.
// If Register is called twice with the same name or if module is nil,
// it panics.
func Register(m of.Module) {
	if m == nil {
		panic("of: Register module is nil")
	}
	if err := _loader.Register(m.Name().ID(), m.Name()); err != nil {
		panic(err)
	}
	err := _loader.Add(m.Name().ID(), m)
	if err != nil {
		panic(err)
	}
}

// AddSource ...
// @param any
// @return error
func AddSource(v any) error {
	return _loader.AddSource(v)
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name) (of.Module, bool) {
	if m, ok := _loader.Get(name.ID()); ok {
		return m, true
	}
	return NoSet, false
}

// GetLoader ...
// @return Loader
func GetLoader() Loader {
	return _loader
}

func MustLoad(name of.Name) of.Module {
	if m, ok := _loader.Get(name.ID()); ok {
		return m
	}
	return NoSet
}

func LoadFromID(id uint64) (of.Module, bool) {
	if m, ok := _loader.Get(id); ok {
		return m, true
	}
	return NoSet, false
}

func MustLoadFromID(id uint64) (of.Module, bool) {
	if m, ok := _loader.Get(id); ok {
		return m, true
	}
	return NoSet, false
}

// LoadModuleNames ...
// @Description: List all registered modules
// @return []of.Name
func LoadModuleNames() []of.Name {
	return _loader.Names()
}
