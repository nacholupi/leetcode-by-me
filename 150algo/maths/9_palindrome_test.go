package maths

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{0, true},
		{123, false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
