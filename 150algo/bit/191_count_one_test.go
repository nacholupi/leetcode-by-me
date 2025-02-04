package bit

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{7, 3},
		{9, 2},
		{15, 4},
		{16, 1},
		{31, 5},
		{32, 1},
		{63, 6},
		{64, 1},
		{127, 7},
		{128, 1},
		{255, 8},
		{256, 1},
	}

	for _, test := range tests {
		result := hammingWeight(test.input)
		if result != test.expected {
			t.Errorf("hammingWeight(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
