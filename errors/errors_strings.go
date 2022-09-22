package errors

type stringErr struct {
	s   string
	err error
}

func (e *stringErr) Error() string {
	return e.String()
}

func (e *stringErr) String() string {
	return e.Error()
}

func (e *stringErr) Message() string {
	return e.String()
}

func (e *stringErr) Unwrap() error {
	return e.err
}

func (e *stringErr) Is(target error) bool {
	if e == target {
		return true
	}
	idx, ok := target.(errMessage)
	if ok && e.Message() == idx.Message() {
		return true
	}
	return false
}

func WrapString(e error, str string) error {
	if e == nil {
		return wrapError(0, str)
	}
	return wrapErrorWithErr(e, 0, str)
}

func WrapStringN(err error, str string) error {
	if err == nil {
		return nil
	}
	return WrapString(err, str)
}
