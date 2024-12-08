package array

import "testing"

func TestCanJumpII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4, 1}, 3},
		{[]int{2, 2, 7, 2, 0, 1, 1 ,1 ,3 ,0 ,1, 1}, 3},
	}

	for _, test := range tests {
		result := canJumpII(test.nums)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
		}
	}
}
