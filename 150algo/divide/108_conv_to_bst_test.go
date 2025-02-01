package divide

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums     []int
		expected *TreeNode
	}{
		{
			nums:     []int{},
			expected: nil,
		},
		{
			nums: []int{1},
			expected: &TreeNode{
				Val: 1,
			},
		},
		{
			nums: []int{1, 2, 3},
			expected: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3}},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}

	for _, test := range tests {
		result := sortedArrayToBST(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, result)
		}
	}
}
