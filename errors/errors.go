package errors

import (
	"fmt"
)

type errorErr struct {
	idx ErrIndex
	err error
	str string
}

type innerError interface {
	Error() string
	String() string
	Unwrap() error
	Index() uint32
	Message() string
	Is(target error) bool
	Has(target error) bool
}

type errIndex interface {
	Index() ErrIndex
}

type errHas interface {
	Has(target error) bool
}

type errMessage interface {
	Message() string
}

func Error(i ErrIndex) error {
	return &indexErr{i: i}
}

func (e *errorErr) Error() string {
	if e.idx != 0 {
		return fmt.Sprintf("[of] %v: %v", e.idx, e.err)
	}
	if e.str != "" {
		return fmt.Sprintf("[of] %v: %v", e.str, e.err)
	}
	return fmt.Sprintf("[of] %v", e.err)
}

func (e *errorErr) String() string {
	return e.Error()
}

func (e *errorErr) Unwrap() error {
	return e.err
}

func (e *errorErr) Index() ErrIndex {
	return e.idx
}

func (e *errorErr) Message() string {
	return e.str
}

func (e *errorErr) Has(target error) bool {
	if e == target {
		return true
	}
	if e.err == nil {
		return false
	}
	has, ok := e.err.(errHas)
	if ok {
		return has.Has(target)
	}
	return false
}

func (e *errorErr) Is(target error) bool {
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

// IndexError
// @param ErrIndex
// @return error
// Decrypted use ErrorIndex instead
func IndexError(i ErrIndex) error {
	return &indexErr{i: i}
}

func ErrorIndex(i ErrIndex) error {
	return &indexErr{i: i}
}

// StringError
// @param string
// @return error
// Decrypted use ErrorString instead
func StringError(s string) error {
	return &stringErr{s: s}
}

func ErrorString(s string) error {
	return &stringErr{s: s}
}

func MessageIs(str string, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errMessage)
	if !ok {
		return false
	}
	return str == e.Message()
}

func IndexIs(i ErrIndex, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errIndex)
	if !ok {
		return false
	}
	return ErrIndex(i) == e.Index()
}
