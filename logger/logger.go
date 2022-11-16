package logger

import (
	"bufio"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

var Default = slog.New(slog.NewJSONHandler(os.Stdout))

type FileLogger struct {
	*slog.Logger
	buf   *bufio.Writer
	close func() error
}

func NewFileLogger(path string) (*FileLogger, error) {
	dir, _ := filepath.Split(path)
	if dir != "" {
		_ = os.MkdirAll(dir, 0755)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	fl := new(FileLogger)
	fl.buf = bufio.NewWriter(file)
	fl.close = file.Close
	fl.Logger = slog.New(slog.NewJSONHandler(fl.buf))
	return fl, nil
}

func (l *FileLogger) Flush() error {
	return l.buf.Flush()
}

func (l *FileLogger) Close() error {
	if l.close != nil {
		err := l.buf.Flush()
		if err != nil {
			return err
		}
		return l.close()
	}
	return nil
}
