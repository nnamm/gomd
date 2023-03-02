package main

import (
	"testing"

	"github.com/nnamm/gomd/clock"
)

func TestSetDirName(t *testing.T) {
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
			s := &source{}
			got := s.setDirName(tt.in)

			if tt.expectErr {
				if got == nil {
					t.Errorf("expectErr: %t, but actual: %v", tt.expectErr, got)
				}
			}

			if s.DirName != tt.want {
				t.Errorf("value to be set: %s, but actual: %s", tt.want, s.DirName)
			}
		})
	}
}

func TestGetCurrentDate(t *testing.T) {
	cases := map[string]struct {
		in   clock.Clocker
		want string
	}{
		"NormalDate1": {clock.FixedClocker{}, "230210"},
	}

	for _, tt := range cases {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			s := &source{}
			d := date{clocker: tt.in}
			s.getCurrentDate(d)

			if s.CurrentDate != tt.want {
				t.Errorf("value to be set: %s, but actual: %s", tt.want, s.CurrentDate)
			}
		})
	}
}
