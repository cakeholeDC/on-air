package logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var TEST_LOG_FILE string

func setupTestEnv(t *testing.T) {
	t.Helper()
	initEnv := os.Getenv("ONAIR_LOG_FILEPATH") // This may not be necessary

	file, err := os.CreateTemp("", "onair-test-log-*.log")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	TEST_LOG_FILE = file.Name()
	defer os.Remove(TEST_LOG_FILE) // Clean up the temporary file

	os.Setenv("ONAIR_LOG_FILEPATH", TEST_LOG_FILE)
	t.Cleanup(func() {
		os.Unsetenv("ONAIR_LOG_FILEPATH")
		os.Remove(TEST_LOG_FILE)
		os.Setenv("ONAIR_LOG_FILEPATH", initEnv) // This may not be necessary.
	})
}

func TestReadEnvLogPath(t *testing.T) {
	setupTestEnv(t)
	// Tests that the log file path is read correctly from the environment variable
	envLog := readEnvLogPath()
	assert.Equal(t, TEST_LOG_FILE, envLog, fmt.Sprintf("Expected config file path to be '%s'", TEST_LOG_FILE))
}

func TestLoggerWritesToLogFile(t *testing.T) {
	setupTestEnv(t)

	// Re-initialize logger to pick up new file path after env is set
	ResetLoggerForTest()
	log := New("testmodule")

	// Write a log entry
	testMsg := "logger_test - This is a test log entry"
	log.Info(testMsg)

	// Wait briefly to ensure log is written
	time.Sleep(100 * time.Millisecond)

	// Open and read the log file
	f, err := os.Open(readEnvLogPath())
	if err != nil {
		t.Fatalf("failed to open log file: %v", err)
	}
	defer f.Close()

	found := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[testmodule]") && strings.Contains(line, testMsg) {
			found = true

			break
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading log file: %v", err)
	}

	if !found {
		t.Errorf("log entry not found in log file")
	}
}
