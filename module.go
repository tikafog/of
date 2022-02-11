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
func Register(name module.Name, m module.Module) {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m == nil {
		panic("of: Register module is nil")
	}
	if name >= module.NameMax {
		panic("of: Register called over module max " + name.String())
	}
	modules[name] = m
}

func Load() []module.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	return modules[:]
}
