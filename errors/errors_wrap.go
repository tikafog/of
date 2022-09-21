package errors

import (
	"errors"
	"fmt"
)

func New(str string) error {
	return errors.New(str)
}

func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func UnwrapError(err error) error {
	return errors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func wrapError(i ErrIndex, str string) error {
	return &errorErr{
		idx: i,
		str: str,
	}
}
func wrapErrorWithErr(parent error, i ErrIndex, str string) error {
	return &errorErr{
		idx: i,
		str: str,
		err: parent,
	}
}
