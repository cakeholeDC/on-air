package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Logger is the exported logger instance
var Logger *slog.Logger

type customHandler struct {
	module string
}

func (h *customHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *customHandler) Handle(_ context.Context, r slog.Record) error {
	timestamp := r.Time.Format("2006/01/02 15:04:05")
	level := r.Level.String()
	fmt.Fprintf(os.Stdout, "%s [%s] %s %s\n", timestamp, h.module, level, r.Message)
	return nil
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return h
}

// Init initializes the logger with a specified module name
func Init(module string, filename string) {
	handler := &customHandler{module: module}
	Logger = slog.New(handler)
}

func New(module string) *slog.Logger {
	return slog.New(&customHandler{module: module})
}
