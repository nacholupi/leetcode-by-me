package twopointers

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, test := range tests {
		result := isSubsequence(test.s, test.t)
		if result != test.expected {
			t.Errorf("For s: %s and t: %s, expected %v, but got %v", test.s, test.t, test.expected, result)
		}
	}
}
