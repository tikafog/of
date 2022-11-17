package logger

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

const (
	ofLoggerFile = "OF_LOGGER_FILE"
)

var (
	writer         io.Writer = os.Stderr
	WipeData       bool      = false
	WipeDataLength int       = 1024
)

type Logger interface {
	Handler() slog.Handler
	Context() context.Context
	With(args ...any) *slog.Logger
	WithGroup(name string) *slog.Logger
	WithContext(ctx context.Context) *slog.Logger
	Enabled(level slog.Level) bool
	Log(level slog.Level, msg string, args ...any)
	LogDepth(calldepth int, level slog.Level, msg string, args ...any)
	LogAttrs(level slog.Level, msg string, attrs ...slog.Attr)
	LogAttrsDepth(calldepth int, level slog.Level, msg string, attrs ...slog.Attr)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, err error, args ...any)
}

func init() {
	env, exist := os.LookupEnv(ofLoggerFile)
	if !exist {
		return
	}
	err := openWriter(env)
	if err != nil {
		return
	}
	defaultLogger = Default()
}

func openWriter(path string) error {
	dir, _ := filepath.Split(path)
	if dir != "" {
		_ = os.MkdirAll(dir, 0755)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	writer = file
	return nil
}

func Default() *slog.Logger {
	return slog.New(slog.NewJSONHandler(writer))
}

func ToFile(path string) (*slog.Logger, error) {
	dir, _ := filepath.Split(path)
	if dir != "" {
		_ = os.MkdirAll(dir, 0755)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return slog.New(slog.NewJSONHandler(file)), nil
}

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
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

func wipeBytes(bytes []byte) []byte {
	if !WipeData || len(bytes) <= WipeDataLength {
		return bytes
	}
	return bytes[:WipeDataLength]
}

func wipeString(str string) string {
	if !WipeData || len(str) <= WipeDataLength {
		return str
	}
	return str[:WipeDataLength]
}
