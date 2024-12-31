package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"", true},
		{"[", false},
		{"{[()()]}", true},
		{"{[(])}", false},
	}

	for _, test := range tests {
		result := IsValid(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}
