package of

import (
	"sync"

	"github.com/tikalink/of/module"
)

var (
	modulesMu sync.RWMutex
	modules   map[module.Name]module.Module
)

// Register makes a database module available by the provided name.
// If Register is called twice with the same name or if module is nil,
// it panics.
func Register(m module.Module) {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m == nil {
		panic("of: Register module is nil")
	}
	if m.Name() >= module.NameMax {
		panic("of: Register called over module max " + m.Name().String())
	}
	modules[m.Name()] = m
}

func Load(name module.Name) module.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	m, ok := modules[name]
	if ok {
		return m
	}
	return module.EmptyModule(name)
}
