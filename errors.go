package of

import (
	"errors"
	"fmt"
)

type err struct {
	idx int
	err error
	str string
}

func (e *err) Error() string {
	return fmt.Sprintf("[err] %v:%v", e.str, e.err)
}

func (e *err) String() string {
	return e.Error()
}

func (e *err) Unwrap() error {
	return e.err
}

func (e *err) Index() int {
	return e.idx
}

func (e *err) Is(target error) bool {
	idx, ok := target.(interface {
		Index() int
	})
	if ok {
		return e.idx == idx.Index()
	}
	return false
}

func wrapError(i int) error {
	return &err{
		idx: i,
		str: errs[i],
	}
}

func WrapIndexError(e error, i int) error {
	if e == nil {
		return nil
	}
	if i < 0 || i > len(errs)-1 {
		return &err{idx: i, err: e, str: "unknown"}
	}
	return &err{idx: i, err: e, str: errs[i]}
}

func WrapError(e error, str string) error {
	if e == nil {
		return nil
	}
	return &err{err: e, str: str}
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

var _ error = &err{}

var errs = [...]string{
	"core is nil",
	"father not found",
	"no data found",
	"channel closed",
}

func Error(i int) error {
	if i < 0 || i > len(errs)-1 {
		return ErrUnknown
	}
	return wrapError(i)
}

var ErrUnknown = &err{idx: 0, err: nil, str: "unknown error"}

//var ErrFatherNotFound = errors.New("father not found")
//var ErrNoDataFound = errors.New("no data found")
//var ErrChannelClosed = errors.New("channel closed")
