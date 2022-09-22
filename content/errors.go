package content

import (
	"github.com/tikafog/of/errors"
)

var err = errors.RegisterModule("Content")

var (
	ErrWrongVersionType = err.NewIndex("wrong version type")
	ErrWrongExtType     = err.NewIndex("wrong ext type")
	ErrWrongMessageType = err.NewIndex("wrong message type")
	ErrWrongContentType = err.NewIndex("wrong content type")
)

func Error(s string) error {
	return err.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return err.Errorf(format, args...)
}
