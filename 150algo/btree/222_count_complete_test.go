package btree

import (
	"testing"
)

func TestCountNodes(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root:     nil,
			expected: 0,
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: 1,
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
			expected: 3,
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
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: 6,
		},
	}

	for _, test := range tests {
		result := countNodes(test.root)
		if result != test.expected {
			t.Errorf("For root %v, expected %d, but got %d", test.root, test.expected, result)
		}
	}
}
