package content

type contentErr struct {
	v string
}

func (c *contentErr) Error() string {
	return c.v
}

func Error(s string) error {
	return &contentErr{v: s}
}
