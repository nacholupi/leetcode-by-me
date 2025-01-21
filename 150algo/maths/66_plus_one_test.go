package maths

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9, 9, 9, 9}, []int{1, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		result := plusOne(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("plusOne(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
