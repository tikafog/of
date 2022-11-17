package logger

import (
	"context"

	"golang.org/x/exp/slog"
)

var defaultLogger Logger = Default()

func Handler() slog.Handler {
	return defaultLogger.Handler()
}

func Context() context.Context {
	return defaultLogger.Context()
}

func With(args ...any) *slog.Logger {
	return defaultLogger.With(args...)
}

func WithGroup(name string) *slog.Logger {
	return defaultLogger.WithGroup(name)
}

func WithContext(ctx context.Context) *slog.Logger {
	return defaultLogger.WithContext(ctx)
}

func Enabled(level slog.Level) bool {
	return defaultLogger.Enabled(level)
}

func Log(level slog.Level, msg string, args ...any) {
	defaultLogger.Log(level, wipeString(msg), args...)
}

func LogDepth(calldepth int, level slog.Level, msg string, args ...any) {
	defaultLogger.LogDepth(calldepth, level, wipeString(msg), args...)
}

func LogAttrs(level slog.Level, msg string, attrs ...slog.Attr) {
	defaultLogger.LogAttrs(level, wipeString(msg), attrs...)
}

func LogAttrsDepth(calldepth int, level slog.Level, msg string, attrs ...slog.Attr) {
	defaultLogger.LogAttrsDepth(calldepth, level, wipeString(msg), attrs...)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debug(wipeString(msg), args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Info(wipeString(msg), args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warn(wipeString(msg), args...)
}

func Error(msg string, err error, args ...any) {
	defaultLogger.Error(wipeString(msg), err, args...)
}
