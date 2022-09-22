package of

import (
	"github.com/tikafog/of/merr"
)

var (
	errors = merr.RegisterModule("Core")
)

var (
	ErrCoreIsNil      = errors.NewIndex("core is nil")
	ErrFatherNotFound = errors.NewIndex("father not found")
	ErrNoDataFound    = errors.NewIndex("no data found")
	ErrChannelClosed  = errors.NewIndex("channel closed")
	ErrSkipOldData    = errors.NewIndex("skip old data")
)

func init() {

	//errors.RegisterErrIndexValue(ErrFatherNotFound, "father not found")
	//errors.RegisterErrIndexValue(ErrNoDataFound, "no data found")
	//errors.RegisterErrIndexValue(ErrChannelClosed, "channel closed")
	//errors.RegisterErrIndexValue(ErrSkipOldData, "skip old data")
	//RegisterErrIndexValue(ErrCoreIsNil, "core is nil")
	//RegisterErrIndexValue(ErrFatherNotFound, "father not found")
	//RegisterErrIndexValue(ErrNoDataFound, "no data found")
	//RegisterErrIndexValue(ErrChannelClosed, "channel closed")
	//RegisterErrIndexValue(ErrSkipOldData, "skip old data")
}

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
