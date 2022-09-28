package dbc

import (
	"github.com/tikafog/errors"
)

var merr = errors.RegisterModule("DBC")

func Error(s string) error {
	return merr.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return merr.Errorf(format, args...)
}
