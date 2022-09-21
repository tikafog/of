package errors

import (
	"fmt"
)

//ErrIndex is used for make error index
type ErrIndex uint32

const (
	ErrIndexCorePrefix    = 0x0001
	ErrIndexContentPrefix = 0x0002
	ErrIndexDBCPrefix     = 0x0003
)

const (
	ErrCoreIsNil      ErrIndex = ErrIndexCorePrefix<<16 | iota
	ErrFatherNotFound ErrIndex = ErrIndexCorePrefix<<16 | iota
	ErrNoDataFound    ErrIndex = ErrIndexCorePrefix<<16 | iota
	ErrChannelClosed  ErrIndex = ErrIndexCorePrefix<<16 | iota
	ErrSkipOldData    ErrIndex = ErrIndexCorePrefix<<16 | iota
)

func init() {
	RegisterErrIndexValue(ErrCoreIsNil, "core is nil")
	RegisterErrIndexValue(ErrFatherNotFound, "father not found")
	RegisterErrIndexValue(ErrNoDataFound, "no data found")
	RegisterErrIndexValue(ErrChannelClosed, "channel closed")
	RegisterErrIndexValue(ErrSkipOldData, "skip old data")
}

var _ErrIndexValue = map[ErrIndex]string{}
var _ErrValueIndex = map[string]ErrIndex{}

var ErrUnknown = &errorErr{idx: 0, err: nil, str: "unknown error"}

func RegisterErrIndexValue(index ErrIndex, str string) {
	_ErrIndexValue[index] = str
	_ErrValueIndex[str] = index
}

func MakeErrIndex(prefix uint32, index uint32) ErrIndex {
	return ErrIndex((prefix << 16) | index)
}

// ParseErrIndex attempts to convert a string to a ErrIndex.
func ParseErrIndex(name string) (ErrIndex, error) {
	if x, ok := _ErrValueIndex[name]; ok {
		return x, nil
	}
	return ErrIndex(0), fmt.Errorf("%s is not a valid Err", name)
}

// String gets the string value of ErrIndex
func (e ErrIndex) String() string {
	if x, ok := _ErrIndexValue[e]; ok {
		return fmt.Sprintf("Module[%v]:%v", e.Module(), x)
	}
	return fmt.Sprintf("Module[%v]:0x%x", e.Module(), uint32(e))
}

var _ErrIndexModules = map[uint32]string{
	ErrIndexCorePrefix:    "Core",
	ErrIndexContentPrefix: "Content",
	ErrIndexDBCPrefix:     "DBC",
}

// Module ...
func (e ErrIndex) Module() string {
	v, ok := _ErrIndexModules[uint32(e>>16)]
	if ok {
		return v
	}
	return "Unknown"
}
