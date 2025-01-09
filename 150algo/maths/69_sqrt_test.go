package maths

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
		{16, 4},
		{25, 5},
		{26, 5},
		{100, 10},
	}

	for _, test := range tests {
		result := mySqrt(test.input)
		if result != test.expected {
			t.Errorf("mySqrt(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
