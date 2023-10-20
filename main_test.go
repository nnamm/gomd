package main

import (
	"testing"
)

func TestIsValidInput(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"77", true},
		{"6", true},
		{"000", true},
		{"999", true},
		{"a", false},
		{"abc", false},
		{"12a", false},
		{"1234", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isInvalidInput(test.input)
			if result != test.expected {
				t.Errorf("isInvalidInput(%s) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}
