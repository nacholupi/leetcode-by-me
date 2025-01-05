package hashmap

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 4, 5}, 3, false},
		{[]int{1, 2, 3, 4, 1}, 4, true},
	}

	for _, test := range tests {
		result := containsNearbyDuplicate(test.nums, test.k)
		if result != test.expected {
			t.Errorf("containsNearbyDuplicate(%v, %d) = %v; expected %v", test.nums, test.k, result, test.expected)
		}
	}
}
