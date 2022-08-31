package of

import (
	"errors"
	"fmt"
)

//Err
//ENUM(core is nil,father not found,no data found,channel closed,skip old data)
type Err int

type errError struct {
	idx int
	err error
	str string
}

type innerError interface {
	Error() string
	String() string
	Unwrap() error
	Index() int
	Message() string
	Is(target error) bool
}

type errIndex interface {
	Index() int
}

type errMessage interface {
	Message() string
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

func (e *errError) Message() string {
	return e.str
}

func (e *errError) Is(target error) bool {
	if e == target {
		return true
	}
	idx, ok := target.(errIndex)
	if ok && e.idx == idx.Index() {
		return true
	}
	msg, ok := target.(errMessage)
	if ok && e.str == msg.Message() {
		return true
	}

	return false
}

func WrapIndexError(e error, i Err) error {
	if e == nil {
		return wrapError(int(i), i.String())
	}
	return wrapErrorWithErr(e, i, i.String())
}

func WrapIndexErrorN(e error, i Err) error {
	if e == nil {
		return nil
	}
	return WrapIndexError(e, i)
}

func WrapError(e error, str string) error {
	if e == nil {
		return wrapError(0, str)
	}
	return wrapErrorWithErr(e, 0, str)
}

func WrapErrorN(err error, str string) error {
	if err == nil {
		return nil
	}
	return WrapError(err, str)
}

func IndexError(i Err) error {
	return wrapError(int(i), i.String())
}

func UnwrapError(err error) error {
	return errors.Unwrap(err)
}

func ErrorIs(err, target error) bool {
	return errors.Is(err, target)
}

func ErrorMessageIs(str string, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errMessage)
	if !ok {
		return false
	}
	return str == e.Message()
}

func ErrorIndexIs(i Err, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errIndex)
	if !ok {
		return false
	}
	return int(i) == e.Index()
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
