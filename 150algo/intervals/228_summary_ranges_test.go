package intervals

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{"1"}},
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{-1}, []string{"-1"}},
		{[]int{1, 3, 5, 7}, []string{"1", "3", "5", "7"}},
	}

	for _, test := range tests {
		result := summaryRanges(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.nums, test.expected, result)
		}
	}
}
