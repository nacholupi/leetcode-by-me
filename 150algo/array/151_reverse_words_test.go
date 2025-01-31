package array

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := reverseWords(test.input)
		if result != test.expected {
			t.Errorf("reverseWords(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
