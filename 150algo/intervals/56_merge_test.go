package intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {0, 4}},
			expected:  [][]int{{0, 4}},
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}},
			expected:  [][]int{{1, 4}},
		},
		{
			intervals: [][]int{},
			expected:  [][]int{},
		},
	}

	for _, test := range tests {
		result := merge(test.intervals)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For intervals %v, expected %v, but got %v", test.intervals, test.expected, result)
		}
	}
}
