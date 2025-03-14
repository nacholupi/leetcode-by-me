package btree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		expected *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			expected: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			expected: &TreeNode{
				Val: -1,
			},
		},
	}

	for _, test := range tests {
		result := buildTree(test.preorder, test.inorder)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("buildTree(%v, %v) = %v; expected %v", test.preorder, test.inorder, result, test.expected)
		}
	}
}
