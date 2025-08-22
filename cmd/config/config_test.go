package config

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateConfig(t *testing.T) {
	// TODO: FIX THIS TEST
	t.Skip("skipping test")
	cmd := ConfigCmd
	args := []string{"--create"}
	cmd.SetArgs(args)

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// NOTE: It's important that we call the 'Execute()' function defined
	// by root.go , rather than 'cmd.Execute()'
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	actual_output := string(out)
	expected := "\033[1monair\033[0m is a command line tool"

	if !strings.Contains(actual_output, expected) {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, actual_output)
	}

	assert.NoError(t, err)
}

// func Test_ConfigList(t *testing.T) {
// 	cmd := rootCmd.Root()
// 	args := []string{"config", "--list"}
// 	cmd.SetArgs(args)

// 	b := bytes.NewBufferString("")
// 	cmd.SetOut(b)

// 	// NOTE: It's important that we call the 'Execute()' function defined
// 	// by root.go , rather than 'cmd.Execute()'
// 	Execute()
// 	out, err := io.ReadAll(b)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	actual_output := string(out)
// 	expected := "\033[1monair\033[0m is a command line tool"

// 	if !strings.Contains(actual_output, expected) {
// 		t.Fatalf("expected \"%s\" got \"%s\"", expected, actual_output)
// 	}

// 	assert.NoError(t, err)
// }
