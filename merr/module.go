package merr

import (
	"errors"
)

type ModuleError interface {
	Name() string
	Index() uint32
	Errors() int
	New(str string) Index
	Errorf(format string, args ...any) Index
	Wrap(err error, s string) Index
	WrapIndex(err error, index Index) Index
	IndexString(index Index) string
	IndexError(index Index) error
	Is(err, target error) bool
	Has(err error) bool
	Unwrap(err error) error
	As(err error, target any) bool
}

type moduleError struct {
	idx    uint32
	name   string
	count  uint32
	errors []error
}

func (m *moduleError) WrapIndex(err error, index Index) Index {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	werr := WrapIndex(err, index)
	m.errors = append(m.errors, werr)
	return idx
}

func (m *moduleError) IndexString(index Index) string {
	//fmt.Sprintf("Module[%v]:%v", index.Name(), err.Error())
	return m.IndexError(index).Error()
}

func (m *moduleError) Name() string {
	return m.name
}

func (m *moduleError) Index() uint32 {
	return m.idx
}

func (m *moduleError) Errors() int {
	return len(m.errors)
}

func (m *moduleError) New(str string) Index {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	err := New(str)
	m.errors = append(m.errors, err)
	return idx
}

func (m *moduleError) Wrap(err error, s string) Index {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	werr := WrapString(err, s)
	m.errors = append(m.errors, werr)
	return idx
}

func (m *moduleError) Errorf(format string, args ...any) Index {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	err := Errorf(format, args...)
	m.errors = append(m.errors, err)
	return idx
}

func (m *moduleError) getErrorIndex(index Index) uint32 {
	idx := (uint32(index) - 1) ^ m.idx<<16
	if idx >= m.count {
		return 0
	}
	return idx
}

func (m *moduleError) IndexError(index Index) error {
	return m.errors[m.getErrorIndex(index)]
}

func (m *moduleError) Is(err, target error) bool {
	return errors.Is(err, target)
}

func (m *moduleError) Unwrap(err error) error {
	return errors.Unwrap(err)
}

func (m *moduleError) As(err error, target any) bool {
	return errors.As(err, target)
}

func (m *moduleError) Has(err error) bool {
	for _, merr := range m.errors {
		if Is(err, merr) {
			return true
		}
	}
	return false
}

func newModuleWithIndex(str string, idx uint32) ModuleError {
	return &moduleError{
		idx:    idx,
		name:   str,
		count:  0,
		errors: nil,
	}
}

func NewModule(str string) ModuleError {
	return newModuleWithIndex(str, getModuleIndex())
}
