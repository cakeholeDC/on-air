package logger

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/natefinch/lumberjack"
)

// Logger is the exported logger instance
var Logger *slog.Logger

var LOG_FILE_PATH string = "onair.log"

// used as the in memory log file
var logWriter *lumberjack.Logger

type customHandler struct {
	module string
}

func (h *customHandler) Enabled(_ context.Context, level slog.Level) bool {
	return true
}

func (h *customHandler) Handle(_ context.Context, r slog.Record) error {
	timestamp := r.Time.Format("2006/01/02 15:04:05")
	level := r.Level.String()
	fmt.Fprintf(logWriter, "%s [%s] %s %s\n", timestamp, h.module, level, r.Message)
	return nil
}

func (h *customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *customHandler) WithGroup(name string) slog.Handler {
	return h
}

// init sets up the rotating log writer
func init() {
	logWriter = &lumberjack.Logger{
		// TODO: make this configurable
		Filename:   LOG_FILE_PATH,
		MaxSize:    1, // megabytes
		MaxBackups: 5, // number of backups
		MaxAge:     7, // days, 0 means no removal based on age
		Compress:   false,
	}
}

// Init initializes the logger with a specified module name
func Init(module string) {
	handler := &customHandler{module: module}
	Logger = slog.New(handler)
}

// New creates a new logger for a module, always using the log file
func New(module string) *slog.Logger {
	return slog.New(&customHandler{module: module})
}
