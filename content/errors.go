package content

type contentErr struct {
	v string
}

func (e *contentErr) Error() string {
	return e.v
}

func (e *contentErr) Is(target error) bool {
	if e == target {
		return true
	}
	msg, ok := target.(*contentErr)
	if ok {
		return e.v == msg.Error()
	}
	return false
}

func Error(s string) error {
	return &contentErr{v: s}
}
