package of

import (
	"errors"
	"fmt"
)

//Err
//ENUM(core is nil,father not found,no data found,channel closed)
type Err int

type errError struct {
	idx int
	err error
	str string
}

type errIndex interface {
	Index() int
}

var ErrUnknown = &errError{idx: 0, err: nil, str: "unknown error"}

func (e *errError) Error() string {
	return fmt.Sprintf("[of] %v: %v", e.str, e.err)
}

func (e *errError) String() string {
	return e.Error()
}

func (e *errError) Unwrap() error {
	return e.err
}

func (e *errError) Index() int {
	return e.idx
}

func (e *errError) Is(target error) bool {
	if e == target {
		return true
	}
	idx, ok := target.(errIndex)
	if ok {
		return e.idx == idx.Index()
	}
	return false
}

func WrapIndexError(e error, i Err) error {
	if e == nil {
		return wrapError(int(i), i.String())
	}
	return wrapErrorWithErr(e, i, i.String())
}

func WrapError(e error, str string) error {
	if e == nil {
		return wrapError(0, str)
	}
	return wrapErrorWithErr(e, 0, str)
}

func IndexError(i Err) error {
	return wrapError(int(i), i.String())
}

func UnwrapError(err error) error {
	return errors.Unwrap(err)
}

func wrapError(i int, str string) error {
	return &errError{
		idx: i,
		str: str,
	}
}
func wrapErrorWithErr(parent error, i Err, str string) error {
	return &errError{
		idx: int(i),
		str: str,
		err: parent,
	}
}
