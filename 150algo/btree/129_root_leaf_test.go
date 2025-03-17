package btree

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
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
			expected: 25, // 12 + 13
		},
		{
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			expected: 1026, // 495 + 491 + 40
		},
		{
			root: &TreeNode{
				Val: 0,
			},
			expected: 0, // Single node tree
		},
	}

	for _, test := range tests {
		result := sumNumbers(test.root)
		if result != test.expected {
			t.Errorf("For tree %v, expected %d but got %d", test.root, test.expected, result)
		}
	}
}
