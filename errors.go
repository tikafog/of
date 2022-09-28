package of

import (
	"github.com/tikafog/errors"
)

var (
	merr = errors.RegisterModule("Core")
)

var (
	ErrCoreIsNil      = merr.New("core is nil")
	ErrFatherNotFound = merr.New("father not found")
	ErrNoDataFound    = merr.New("no data found")
	ErrChannelClosed  = merr.New("channel closed")
	ErrSkipOldData    = merr.New("skip old data")
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
