package bit

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b, expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"1111", "1111", "11110"},
		{"110", "101", "1011"},
	}

	for _, test := range tests {
		result := addBinary(test.a, test.b)
		if result != test.expected {
			t.Errorf("addBinary(%s, %s) = %s; expected %s", test.a, test.b, result, test.expected)
		}
	}
}
