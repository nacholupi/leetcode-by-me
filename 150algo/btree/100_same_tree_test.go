package btree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: true,
		},
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:        &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: false,
		},
		{
			p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			expected: false,
		},
		{
			p:        nil,
			q:        nil,
			expected: true,
		},
	}

	for _, test := range tests {
		result := isSameTree(test.p, test.q)
		if result != test.expected {
			t.Errorf("isSameTree(%v, %v) = %v; expected %v", test.p, test.q, result, test.expected)
		}
	}
}
