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
