package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerateContent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"6", "006_" + currentDate(), false},
		{"77", "077_" + currentDate(), false},
		{"123", "123_" + currentDate(), false},
		{"000", "000_" + currentDate(), false},
		{"999", "999_" + currentDate(), false},
		{"a", "", true},
		{"abc", "", true},
		{"12a", "", true},
		{"1234", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		err := runTestGenerateContent(t, test)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
	}
}

func currentDate() string {
	t := time.Now()
	year, month, day := t.Date()
	return fmt.Sprintf("%02d%02d%02d", year%100, month, day)
}

func runTestGenerateContent(t *testing.T, test struct {
	input    string
	expected string
	hasError bool
}) error {
	err := GenerateContent(test.input)

	defer os.RemoveAll(test.expected)

	if test.hasError {
		if err == nil {
			t.Errorf("expected an error for input %v", test.input)
			return err
		}
		return nil
	}

	if err != nil {
		t.Errorf("unexpected error for input %v: %v", test.input, err)
		return err
	}

	// check if directory was created
	_, err = os.Stat(test.expected)
	if os.IsNotExist(err) {
		t.Errorf("directory %v was not created", test.expected)
		return err
	}

	// check if markdown file exists
	mdPath := filepath.Join(test.expected, "index.md")
	info, err := os.Stat(mdPath)
	if os.IsNotExist(err) || info.IsDir() {
		t.Errorf("markdown file %v was not created", mdPath)
		return err
	}

	return nil
}
