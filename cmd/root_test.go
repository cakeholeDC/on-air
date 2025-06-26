package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test has little value on it's own, but it's here as reference
// for future tests that may need to test other parts of the CLI.
//
// The boilerplate required may be confusing to some, but using the
// pattern below, one can test other subcommands & arguments
//
// For example, consider:
//
//	args := []string{"config", "--hass-endpoint" "https://local-dsp.virtru.com:8080"}
func Test_GetHelp(t *testing.T) {
	cmd := rootCmd.Root()
	args := []string{"--help"}
	cmd.SetArgs(args)

	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// NOTE: It's important that we call the 'Execute()' function defined
	// by root.go , rather than 'cmd.Execute()'
	Execute()
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
