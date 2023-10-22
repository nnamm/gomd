package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"gomd", "12"}, ""}, // delete generated dir after each testing
		{[]string{"gomd", "abc"}, "invalid argument, expected 1-3 digit number\n"},
		{[]string{"gomd"}, "expected a single argument\n"},
		{[]string{"gomd", "12", "12"}, "expected a single argument\n"},
	}

	for _, test := range tests {
		os.Args = test.input

		// capture stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		main()

		// reset stdout
		os.Stdout = oldStdout
		w.Close()

		var buf bytes.Buffer
		buf.ReadFrom(r)
		output := buf.String()

		if output != test.expected {
			t.Errorf("expected %q but got %q", test.expected, output)
		}
	}
}
