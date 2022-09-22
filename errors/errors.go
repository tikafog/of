package errors

type errorErr struct {
	i   Index
	str string
	err error
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
	Index() Index
}

type errHas interface {
	Has(target error) bool
}

type errMessage interface {
	Message() string
}

func (e *errorErr) Error() string {
	return e.i.String() + ":" + e.str
}

func (e *errorErr) String() string {
	return e.Error()
}

func (e *errorErr) Unwrap() error {
	return e.err
}

func (e *errorErr) Index() Index {
	return e.i
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
	if ok && e.i == idx.Index() {
		return true
	}
	msg, ok := target.(errMessage)
	if ok && e.str == msg.Message() {
		return true
	}
	return false
}

func Error(i Index) error {
	return &indexErr{i: i}
}

// IndexError
// @param Index
// @return error
// Decrypted use ErrorIndex instead
func IndexError(i Index) error {
	return &indexErr{i: i}
}

func ErrorIndex(i Index) error {
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

func IndexIs(i Index, err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(errIndex)
	if !ok {
		return false
	}
	return Index(i) == e.Index()
}
