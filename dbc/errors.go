package dbc

import (
	"github.com/tikafog/of/errors"
)

const ErrIndex = 0x0003

func Error(s string) error {
	return errors.ModuleError(ErrIndex, errors.New(s))
}

func Errorf(format string, args ...interface{}) error {
	return errors.ModuleError(ErrIndex, errors.Errorf(format, args...))
}
