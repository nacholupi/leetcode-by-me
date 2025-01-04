package bfs

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected []int
	}{
		{
			root:     nil,
			expected: []int{},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: []int{1, 3, 5},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			expected: []int{1, 3, 4},
		},
	}

	for _, test := range tests {
		result := rightSideView(test.root)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, result)
		}
	}
}
