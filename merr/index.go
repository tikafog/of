package merr

import (
	"fmt"
)

type indexErr struct {
	i   Index
	err error
}

func (e *indexErr) Error() string {
	return e.String()
}

func (e *indexErr) String() string {
	if e.err != nil {
		return fmt.Sprintf("%v: %v", e.i, e.err)
	}
	return e.i.String()
}

func (e *indexErr) Index() Index {
	return e.i
}

func (e *indexErr) Message() string {
	return e.String()
}

func (e *indexErr) Unwrap() error {
	return e.err
}

func (e *indexErr) Is(target error) bool {
	if e == target {
		return true
	}
	idx, ok := target.(errIndex)
	if ok && e.i == idx.Index() {
		return true
	}
	return false
}

func makeErrIndex(prefix uint32, index uint32) Index {
	return Index((prefix << 16) | index)
}

func WrapIndex(e error, i Index) error {
	if e == nil {
		return &indexErr{i: i}
	}
	return &indexErr{i: i, err: e}
}

func WrapIndexN(e error, i Index) error {
	if e == nil {
		return nil
	}
	return &indexErr{i: i, err: e}
}
