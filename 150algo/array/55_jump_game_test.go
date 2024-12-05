package array

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{0}, true},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 0, 0}, true},
		{[]int{1, 2, 3}, true},
	}

	for _, test := range tests {
		result := canJump(test.nums)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
		}
		result = canJumpRecursive(test.nums)
		if result != test.expected {
			t.Errorf("Recursive: For input %v, expected %v but got %v", test.nums, test.expected, result)
		}
	}
}
