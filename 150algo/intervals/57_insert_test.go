package intervals

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{4, 5},
			expected:    [][]int{{1, 3}, {4, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			expected:    [][]int{{1, 5}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			expected:    [][]int{{1, 7}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 5}, {6, 8}},
		},
	}

	for _, test := range tests {
		result := insert(test.intervals, test.newInterval)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For intervals %v and newInterval %v, expected %v but got %v", test.intervals, test.newInterval, test.expected, result)
		}
	}
}
