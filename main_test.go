package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDirNo(t *testing.T) {
	cases := map[string]struct {
		in        string
		want      string
		expectErr bool
	}{
		"DirNo Length: 1":                  {"0", "000", false},
		"DirNo Length: 2":                  {"00", "000", false},
		"DirNo Length: 3":                  {"000", "000", false},
		"DirNo Length: 4":                  {"0000", "", true},
		"DirNo Length: 3 include alphabet": {"A00", "", true},
		"DirNo Length: 3 include symbol":   {"00*", "", true},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			s := Source{}
			got := s.setDirNo(tt.in)

			fmt.Println(tt.in)

			if tt.expectErr {
				if got == nil {
					t.Errorf("expectErr: %t, but actual: %v", tt.expectErr, got)
				}
			}

			if s.DirNo != tt.want {
				t.Errorf("value to be set: %s, but actual: %s", tt.want, s.DirNo)
			}
		})
	}
}

func TestDateFormat(t *testing.T) {
	cases := map[string]struct {
		in        time.Time
		want      string
		expectErr bool
	}{
		"Today1": {time.Date(2023, time.February, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)), "_230201", false},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.in.String(), func(t *testing.T) {
			t.Parallel()
			got, err := dateFormat()

			if tt.expectErr {
				if err == nil {
					t.Errorf("expectErr: %t, but actual err: %v", tt.expectErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("expectErr: %t, but actual err: %v", tt.expectErr, err)
				}
			}

			if got != tt.want {
				t.Errorf("value to be set: %s, but actual: %s", tt.want, got)
			}
		})
	}
}
