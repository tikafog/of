package of

import (
	"sync"

	"github.com/tikalink/of/module"
)

var (
	modulesMu sync.RWMutex
	modules   [module.NameMax]module.Module
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

func Load() []module.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	ms := make([]module.Module, 0)
	for i := range modules {
		if modules[i] != nil {
			ms = append(ms, modules[i])
		}
	}
	return ms
}
