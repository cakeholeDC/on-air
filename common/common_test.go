package common

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintJSON(t *testing.T) {
	// Redirect output to capture the printed JSON
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test data
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// Call the function
	PrintJSON(data)

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	io.Copy(&buf, r)

	// Expected JSON output
	expected := `{
  "age": 30,
  "name": "John"
}` + "\n"

	if buf.String() != expected {
		t.Errorf("Expected %s but got %s", expected, buf.String())
	}
}

func TestPrintYAML(t *testing.T) {
	// Redirect output to capture the printed YAML
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test data
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// Call the function
	PrintYAML(data)

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read the captured output
	io.Copy(&buf, r)

	// Expected YAML output
	expected := `age: 30
name: John
` + "\n"

	if buf.String() != expected {
		t.Errorf("Expected %s but got %s", expected, buf.String())
	}
}

func TestGetEnv(t *testing.T) {
	// Test the Getenv function
	os.Setenv("TEST_ENV", "test_value")
	defer os.Unsetenv("TEST_ENV") // Clean up after test

	value := GetEnv("TEST_ENV", "default_value")
	assert.Equal(t, "test_value", value, "Expected 'test_value' for TEST_ENV")

	defaultValue := GetEnv("NON_EXISTENT_ENV", "default_value")
	assert.Equal(t, "default_value", defaultValue, "Expected 'default_value' for NON_EXISTENT_ENV")
}

func TestGetTerminalSize(t *testing.T) {
	// Test the getTerminalSize function
	width, height, _ := GetTerminalSize()
	// swallow the error for this test
	assert.Greater(t, width, 0, "Expected terminal width to be greater than 0")
	assert.Greater(t, height, 0, "Expected terminal height to be greater than 0")
}

func TestPrintArtStringIfFits(t *testing.T) {
	// Test the PrintArtStringIfFits function
	// This test will not fail, but it will print the art string if it fits in the terminal width.
	// It is a visual test, so we will not assert anything here.
	content := "Hello, Text Art!"
	PrintArtStringIfFits(content)
}

func TestReadWriteFileBlob(t *testing.T) {
	// Test the ReadFileBlob function
	testFile := "test.txt"
	plainText := "Hello, World!"
	err := WriteFileBlob(testFile, []byte(plainText))
	if err != nil {
		t.Errorf("Error writing file: %v", err)
	}
	blob, err := ReadFileBlob(testFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}
	assert.Equal(t, string(blob), plainText)
	os.Remove(testFile)
}
