package content

import (
	"github.com/tikafog/of/merr"
)

var errors = merr.RegisterModule("Content")

var (
	ErrWrongVersionType = errors.NewIndex("wrong version type")
	ErrWrongExtType     = errors.NewIndex("wrong ext type")
	ErrWrongMessageType = errors.NewIndex("wrong message type")
	ErrWrongContentType = errors.NewIndex("wrong content type")
)

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
