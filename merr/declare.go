package merr

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

//Index is used for make error index
type Index uint32

var (
	modulesMu     sync.RWMutex
	moduleIndex   uint32
	moduleErrors  = make(map[string]ModuleError)
	moduleIndexes = make(map[uint32]string)
)

var (
	ErrModuleIsAlreadyRegistered = NewIndex("module is already registered")
	ErrModuleSupportOverMax      = NewIndex("module support over max")
)

var UnknownError = IndexError(0)
var UnknownModule = registerModuleWithIndex("unknow", 0)

func NewIndex(s string) Index {
	return UnknownModule.NewIndex(s)
}

func getModuleIndex() uint32 {
	return atomic.AddUint32(&moduleIndex, 1)
}

// String gets the string value of Index
func (e Index) String() string {
	//err := Module(e.Name()).IndexString(e)
	err := Module(e.Name()).IndexError(e)
	return fmt.Sprintf("Module[%v]: %v", e.Name(), err.Error())
	//return
}

func Module(name string) ModuleError {
	modulesMu.RLock()
	e, ok := moduleErrors[name]
	modulesMu.RUnlock()
	if ok {
		return e
	}
	return UnknownModule
}

func ModuleName(index uint32) string {
	modulesMu.RLock()
	v, ok := moduleIndexes[index]
	modulesMu.RUnlock()
	if ok {
		return v
	}
	return "unknown"
}

func registerModuleWithIndex(name string, idx uint32) ModuleError {
	name = strings.ToLower(name)
	modulesMu.Lock()
	modulesMu.Unlock()
	if e, ok := moduleErrors[name]; ok {
		return e
	}
	m := newModuleWithIndex(name, idx)
	moduleErrors[name] = m
	moduleIndexes[m.Index()] = m.Name()
	fmt.Println("registered module:", name, "index:", m.Index())
	return moduleErrors[name]
}

func RegisterModule(name string) ModuleError {
	return registerModuleWithIndex(name, getModuleIndex())
}

// Name ...
func (e Index) Name() string {
	return ModuleName(e.ModuleIndex())
}

// ModuleIndex ...
func (e Index) ModuleIndex() uint32 {
	return uint32(e) >> 16
}

func (e Index) Error() string {
	return e.String()
}

// IndexModule ...
func IndexModule(e uint32) uint32 {
	return e << 16
}

func (e Index) Module() ModuleError {
	return Module(e.Name())
}
