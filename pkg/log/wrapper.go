package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type Wrapper interface {
	Debugf(string, ...any)
	Infof(string, ...any)
	Warnf(string, ...any)
	Errorf(string, ...any)
	Fatalf(string, ...any)
}

func newWrapper() Wrapper {
	return &wrapper{}
}

type wrapper struct {
}

func (w *wrapper) logf(lvl slog.Level, msg string, args ...any) {
	slog.Log(context.Background(), lvl, fmt.Sprintf(msg, args))
}

func (w *wrapper) Debugf(msg string, args ...any) {
	w.logf(slog.LevelDebug, msg, args...)
}

func (w *wrapper) Infof(msg string, args ...any) {
	w.logf(slog.LevelInfo, msg, args...)
}

func (w *wrapper) Warnf(msg string, args ...any) {
	w.logf(slog.LevelWarn, msg, args...)
}

func (w *wrapper) Errorf(msg string, args ...any) {
	w.logf(slog.LevelError, msg, args...)
}

func (w *wrapper) Fatalf(msg string, args ...any) {
	w.logf(slog.LevelError, msg, args...)

	os.Exit(1)
}
