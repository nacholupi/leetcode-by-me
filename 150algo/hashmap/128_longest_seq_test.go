package hashmap

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}

	for _, test := range tests {
		result := longestConsecutive(test.nums)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
