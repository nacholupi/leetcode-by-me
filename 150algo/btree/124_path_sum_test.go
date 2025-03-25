package btree

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "single node",
			root:     &TreeNode{Val: 5},
			expected: 5,
		},
		{
			name: "two nodes",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			expected: 3,
		},
		{
			name: "three nodes",
			root: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 10},
			},
			expected: 22,
		},
		{
			name: "complex tree",
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 20},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: -25, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				},
			},
			expected: 42,
		},
		{
			name: "negative values",
			root: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: -20},
				Right: &TreeNode{
					Val:  -30,
					Left: &TreeNode{Val: -5},
				},
			},
			expected: -5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxPathSum(test.root)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
