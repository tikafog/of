package instruct

import (
	"github.com/tikafog/errors"
)

var merr = errors.RegisterModule("metaInstruct")

var (
	ErrWrongVersionType  = merr.New("wrong version type")
	ErrWrongExtType      = merr.New("wrong ext type")
	ErrWrongMessageType  = merr.New("wrong message type")
	ErrWrongInstructType = merr.New("wrong instruct type")
)

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
