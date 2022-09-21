package content

import (
	"github.com/tikafog/of/errors"
)

const ErrIndex = 0x0002

const (
	ErrContent          = ErrIndex<<16 | iota
	ErrWrongVersionType = ErrIndex<<16 | iota
	ErrWrongExtType     = ErrIndex<<16 | iota
	ErrWrongMessageType = ErrIndex<<16 | iota
	ErrWrongContentType = ErrIndex<<16 | iota
)

func init() {
	errors.RegisterErrIndexValue(ErrWrongVersionType, "wrong version type")
	errors.RegisterErrIndexValue(ErrWrongExtType, "wrong ext type")
	errors.RegisterErrIndexValue(ErrWrongMessageType, "wrong message type")
	errors.RegisterErrIndexValue(ErrWrongContentType, "wrong content type")
}

func Error(s string) error {
	return errors.ModuleError(ErrIndex, errors.New(s))
}

func Errorf(format string, args ...interface{}) error {
	return errors.ModuleError(ErrIndex, errors.Errorf(format, args...))
}
