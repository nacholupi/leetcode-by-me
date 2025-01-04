package bfs

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			root:     nil,
			expected: [][]int{},
		},
		{
			root:     &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: [][]int{{1}, {2, 3}},
		},
		//[1,2,null,3,null,4,null,5]
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			expected: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6}},
		},
	}

	for _, test := range tests {
		result := levelOrder(test.root)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, result)
		}
	}
}
