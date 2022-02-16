package module

import (
	"sync"

	"github.com/tikalink/of"
	"github.com/tikalink/of/option"
)

var (
	modulesMu sync.RWMutex
	modules   = map[of.Name]Loader{}
)

type Loader interface {
	of.Module
	Option(option.Option) of.Module
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
	if m.Name() >= of.NameMax {
		panic("of: Register called over module max " + m.Name().String())
	}
	modules[m.Name()] = m
}

// Load ...
// @Description: Load module
// @param module.Name
// @return module.Module
func Load(name of.Name, op option.Option) of.Module {
	modulesMu.Lock()
	defer modulesMu.Unlock()
	m, ok := modules[name]
	if ok {
		return m.Option(op)
	}
	return EmptyModule(name)
}
