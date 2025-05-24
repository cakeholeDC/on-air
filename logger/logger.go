package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Logger is the exported logger instance
var Logger *slog.Logger

var LOG_FILE_PATH string = "onair.log"

// used as the in memory log file
var logFile *os.File

type customHandler struct {
	module string
	file   *os.File
}

func (h *customHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *customHandler) Handle(_ context.Context, r slog.Record) error {
	timestamp := r.Time.Format("2006/01/02 15:04:05")
	level := r.Level.String()
	// Write to file if set, else fallback to stdout
	out := os.Stdout
	if h.file != nil {
		out = h.file
	}
	fmt.Fprintf(out, "%s [%s] %s %s\n", timestamp, h.module, level, r.Message)
	return nil
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return h
}

// init opens the log file once when the package is loaded
func init() {
	var err error
	logFile, err = os.OpenFile(LOG_FILE_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("failed to open log file: %v", err))
	}
}

// Init initializes the logger with a specified module name
func Init(module string) {
	handler := &customHandler{module: module, file: logFile}
	Logger = slog.New(handler)
}

// New creates a new logger for a module, always using the log file
func New(module string) *slog.Logger {
	return slog.New(&customHandler{module: module, file: logFile})
}
