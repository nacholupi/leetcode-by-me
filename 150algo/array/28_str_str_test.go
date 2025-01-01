package array

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"a", "a", 0},
		{"mississippi", "issip", 4},
		{"mississippi", "sippia", -1},
	}

	for _, test := range tests {
		result := strStr(test.haystack, test.needle)
		if result != test.expected {
			t.Errorf("strStr(%q, %q) = %d; expected %d", test.haystack, test.needle, result, test.expected)
		}
	}
}
