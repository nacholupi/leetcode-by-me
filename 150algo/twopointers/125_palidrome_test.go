package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abc", false},
		{"abba", true},
		{"abcd", false},
		{"A man, a plan, a canal, Panama", true},
		{"No 'x' in Nixon", true},
		{"Was it a car or a cat I saw?", true},
		{"Madam, in Eden, I'm Adam", true},
		{"Able was I ere I saw Elba", true},
		{"race a car", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
		}
	}
}
