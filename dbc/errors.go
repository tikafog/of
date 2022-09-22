package dbc

import (
	"github.com/tikafog/of/merr"
)

//var errors = merr.RegisterModule("DBC")

func Error(s string) error {
	return merr.New(s)
}

func Errorf(format string, args ...interface{}) error {
	return merr.Errorf(format, args...)
}
