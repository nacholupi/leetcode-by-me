package main

import "testing"

func TestPartitionString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abac", 2},
		{"world", 1},
		{"aabbcc", 4},
		{"abcabc", 2},
		{"", 0},
	}

	for _, test := range tests {
		result := partitionString(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
