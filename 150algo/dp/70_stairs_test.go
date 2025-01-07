package dp

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
	}
	for _, test := range tests {
		if output := climbStairs(test.n); output != test.expected {
			t.Error("Test Failed: {} expected, received: {}", test.expected, output)
		}
	}
}
