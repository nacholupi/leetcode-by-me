package sliding

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"a", "b", ""},
		{"ab", "a", "a"},
		{"ab", "b", "b"},
		{"ab", "ab", "ab"},
		{"bba", "ab", "ba"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
	}

	for _, test := range tests {
		result := minWindow(test.s, test.t)
		if result != test.expected {
			t.Errorf("minWindow(%q, %q) = %q; expected %q", test.s, test.t, result, test.expected)
		}
	}
}
