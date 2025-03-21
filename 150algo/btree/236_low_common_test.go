package btree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	// Helper function to create a new TreeNode
	newNode := func(val int) *TreeNode {
		return &TreeNode{Val: val}
	}

	// Construct the binary tree
	root := newNode(3)
	root.Left = newNode(5)
	root.Right = newNode(1)
	root.Left.Left = newNode(6)
	root.Left.Right = newNode(2)
	root.Right.Left = newNode(0)
	root.Right.Right = newNode(8)
	root.Left.Right.Left = newNode(7)
	root.Left.Right.Right = newNode(4)

	tests := []struct {
		p, q     *TreeNode
		expected *TreeNode
	}{
		{root.Left, root.Right, root},                 // LCA of 5 and 1 is 3
		{root.Left, root.Left.Right.Right, root.Left}, // LCA of 5 and 4 is 5
		{root.Left.Left, root.Left.Right, root.Left},  // LCA of 6 and 2 is 5
	}

	for _, test := range tests {
		result := lowestCommonAncestor(root, test.p, test.q)
		if result != test.expected {
			t.Errorf("lowestCommonAncestor(%v, %v) = %v; expected %v", test.p.Val, test.q.Val, result.Val, test.expected.Val)
		}
	}
}
