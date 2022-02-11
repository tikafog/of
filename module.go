package of

import (
	"sync"

	"github.com/tikalink/of/module"
)

var (
	modulesMu sync.RWMutex
	modules   = make(map[string]module.Module)
)

// Register makes a database module available by the provided name.
// If Register is called twice with the same name or if module is nil,
// it panics.
func Register(name string, m module.Module) {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	if m == nil {
		panic("of: Register module is nil")
	}
	if _, dup := modules[name]; dup {
		panic("of: Register called twice for module " + name)
	}
	modules[name] = m
}
