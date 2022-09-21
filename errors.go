package of

import (
	"github.com/tikafog/of/errors"
)

const (
	ErrIndex = 0x0001
)

const (
	ErrCore           = ErrIndex<<16 | iota
	ErrCoreIsNil      = ErrIndex<<16 | iota
	ErrFatherNotFound = ErrIndex<<16 | iota
	ErrNoDataFound    = ErrIndex<<16 | iota
	ErrChannelClosed  = ErrIndex<<16 | iota
	ErrSkipOldData    = ErrIndex<<16 | iota
)

func init() {
	errors.RegisterErrIndexValue(ErrCoreIsNil, "core is nil")
	errors.RegisterErrIndexValue(ErrFatherNotFound, "father not found")
	errors.RegisterErrIndexValue(ErrNoDataFound, "no data found")
	errors.RegisterErrIndexValue(ErrChannelClosed, "channel closed")
	errors.RegisterErrIndexValue(ErrSkipOldData, "skip old data")
}

func Error(s string) error {
	return errors.ModuleError(ErrIndex, errors.New(s))
}

func Errorf(format string, args ...interface{}) error {
	return errors.ModuleError(ErrIndex, errors.Errorf(format, args...))
}
