package hashmap

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"", "", true},
		{"ab", "aa", false},
		{"abc", "def", true},
	}

	for _, test := range tests {
		result := isIsomorphic(test.s, test.t)
		if result != test.expected {
			t.Errorf("isIsomorphic(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
