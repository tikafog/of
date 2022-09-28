package content

import (
	"github.com/tikafog/errors"
)

var merr = errors.RegisterModule("content")

var (
	ErrWrongVersionType = merr.New("wrong version type")
	ErrWrongExtType     = merr.New("wrong ext type")
	ErrWrongMessageType = merr.New("wrong message type")
	ErrWrongContentType = merr.New("wrong content type")
)

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
