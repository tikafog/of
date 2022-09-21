package errors

type indexErr struct {
	i ErrIndex
}

func (e *indexErr) Error() string {
	return e.String()
}

func (e *indexErr) String() string {
	return e.Error()
}

func (e *indexErr) Index() ErrIndex {
	return e.i
}

func (e *indexErr) Message() string {
	return e.String()
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

func WrapIndex(e error, i ErrIndex) error {
	if e == nil {
		return wrapError(i, i.String())
	}
	return wrapErrorWithErr(e, i, i.String())
}

func WrapIndexN(e error, i ErrIndex) error {
	if e == nil {
		return nil
	}
	return WrapIndex(e, i)
}
