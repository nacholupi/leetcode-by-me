package bst

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		k        int
		expected int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			k:        1,
			expected: 1,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			k:        3,
			expected: 3,
		},
	}

	for _, test := range tests {
		result := kthSmallest(test.root, test.k)
		if result != test.expected {
			t.Errorf("For root %v and k %d, expected %d but got %d", test.root, test.k, test.expected, result)
		}
	}
}
