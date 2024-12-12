package array

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4}, expected: []int{24, 12, 8, 6}},
		{input: []int{5, 6, 7, 8}, expected: []int{336, 280, 240, 210}},
		{input: []int{1, 1, 1, 1}, expected: []int{1, 1, 1, 1}},
		{input: []int{2, 3, 4, 5}, expected: []int{60, 40, 30, 24}},
		{input: []int{0, 1, 2, 3}, expected: []int{6, 0, 0, 0}},
	}

	for _, test := range tests {
		result := productExceptSelf(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
