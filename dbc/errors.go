package dbc

import (
	"github.com/tikafog/of/errors"
)

var err = errors.RegisterModule("DBC")

func Error(s string) error {
	return err.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return err.Errorf(format, args...)
}
