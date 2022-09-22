package of

var (
	err = merr.RegisterModule("Core")
)

var (
	ErrCoreIsNil      = err.NewIndex("core is nil")
	ErrFatherNotFound = err.NewIndex("father not found")
	ErrNoDataFound    = err.NewIndex("no data found")
	ErrChannelClosed  = err.NewIndex("channel closed")
	ErrSkipOldData    = err.NewIndex("skip old data")
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
	return err.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return err.Errorf(format, args...)
}
