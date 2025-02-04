package bit

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{0, 1, 0}, 1},
		{[]int{2, 3, 2, 4, 4}, 3},
	}

	for _, test := range tests {
		result := singleNumber(test.nums)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.nums, test.expected, result)
		}
	}
}
