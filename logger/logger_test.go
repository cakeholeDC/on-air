package logger

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLoggerWritesToLogFile(t *testing.T) {
	// Remove log file if it exists
	os.Remove(LOG_FILE_PATH)

	// Initialize logger
	Init("testmodule")

	// Write a log entry
	testMsg := "This is a test log entry"
	Logger.Info(testMsg)

	// Wait briefly to ensure log is written
	time.Sleep(100 * time.Millisecond)

	// Open the log file
	f, err := os.Open(LOG_FILE_PATH)
	if err != nil {
		t.Fatalf("failed to open log file: %v", err)
	}
	defer f.Close()

	// Read the log file
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
