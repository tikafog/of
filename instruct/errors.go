package instruct

import (
	"github.com/tikafog/of/merr"
)

var errors = merr.RegisterModule("metaInstruct")

var (
	ErrWrongVersionType  = errors.New("wrong version type")
	ErrWrongExtType      = errors.New("wrong ext type")
	ErrWrongMessageType  = errors.New("wrong message type")
	ErrWrongInstructType = errors.New("wrong instruct type")
)

func Error(s string) error {
	return errors.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}
