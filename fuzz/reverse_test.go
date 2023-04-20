package main

import (
	"testing"
	"unicode/utf8"
)

// The unit test for declared test cases.
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev, revErr := Reverse(tc.in)
		if revErr != nil {
			// return
			t.Skip()
		}

		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}

}

// Fuzz test.
// The function begins with FuzzXxx instead of TestXxx, and takes *testing.F instead of *testing.T
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}

	for _, tc := range testcases {
		// Use f.Add to provide a seed corpus
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		if revErr != nil {
			// return
			t.Skip()
		}

		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil {
			// return
			t.Skip()
		}

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(orig),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev))

		// The string orig should be equal to string doubleRev. otherwise, test fails.
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}

		// The string `orig` and `rev` are should be passed validation
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
