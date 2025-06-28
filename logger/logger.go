package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/natefinch/lumberjack"
)

// Logger is the exported logger instance
var Logger *slog.Logger

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
		Filename:   readEnvLogPath(),
		MaxSize:    1,     // megabytes
		MaxBackups: 5,     // number of backups
		MaxAge:     7,     // days, 0 means no removal based on age
		Compress:   false, // compress using gzip. The default is not to perform compression.
	}
}

// New creates a new logger for a module, always using the log file
func New(module string) *slog.Logger {
	return slog.New(&customHandler{module: module})
}

func readEnvLogPath() string {
	// Read the enviornment variable for the log file path
	envLogPath := os.Getenv("ONAIR_LOG_FILEPATH")

	// If the environment variable is not set, use the default path
	if envLogPath == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			slog.Error(fmt.Sprintf("could not determine user home directory: %s", err))
			return ""
		}
		envLogPath = fmt.Sprintf("%s/.config/onair/onair.log", userHomeDir)
	}
	// return the log file path
	return envLogPath
}

// ResetLoggerForTest resets the logger for testing purposes
func ResetLoggerForTest() {
	logWriter = &lumberjack.Logger{
		Filename:   readEnvLogPath(),
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   false,
	}
}
