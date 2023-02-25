package main

import "testing"

func TestSource(t *testing.T) {
	cases := map[string]struct {
		dirNo   string
		wantErr bool // true is a case in which an error occurs
	}{
		"Length: 3": {"000", false},
	}

	for _, tt := range cases {
		t.Run(tt.dirNo, func(t *testing.T) {
			s := Source{}
			got := s.setDirFormat(tt.dirNo)

			if tt.wantErr {
				if got == nil {
					t.Errorf("")
				}
			} else {
				if tt.dirNo != s.DirNo {
					t.Errorf("value to be set: %s, but actural: %s", tt.dirNo, s.DirNo)
				}
			}
		})
	}
	//	t.Run("Length:3", func(t *testing.T) {
	//		s := Source{}
	//		got := s.setDirFormat("000")
	//
	//		if got != nil {
	//			t.Errorf("want nil, but got err: %v", got)
	//		}
	//	})
	//
	//t.Run("Length:2", func(t *testing.T) {
	//	s := Source{}
	//	got := s.setDirFormat("00")
	//
	//	if got != nil {
	//		t.Errorf("want nil, but got err: %v", got)
	//	}
	//})
	//
	//t.Run("Length:4", func(t *testing.T) {
	//	s := Source{}
	//	got := s.setDirFormat("0000")
	//
	//	if got != nil {
	//		t.Errorf("want nil, but got err: %v", got)
	//	}
	//})
	//
	//t.Run("Length:3 but NOT all are numbers", func(t *testing.T) {
	//	s := Source{}
	//	got := s.setDirFormat("0A0")
	//
	//	if got != nil {
	//		t.Errorf("want nil, but got err: %v", got)
	//	}
	//})
}
