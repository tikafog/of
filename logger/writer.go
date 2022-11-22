package logger

import (
	"io"
)

type writer struct {
	io.Writer
}

func newWriter(w io.Writer) *writer {
	return &writer{
		Writer: w,
	}
}

func (w *writer) SetOutput(out io.Writer) {
	w.Writer = out
}
