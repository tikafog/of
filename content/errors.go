package content

import (
	"github.com/tikafog/of/merr"
)

var errors = merr.RegisterModule("Content")

var (
	ErrWrongVersionType = errors.New("wrong version type")
	ErrWrongExtType     = errors.New("wrong ext type")
	ErrWrongMessageType = errors.New("wrong message type")
	ErrWrongContentType = errors.New("wrong content type")
)

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
