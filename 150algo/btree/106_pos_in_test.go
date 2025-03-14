package btree

import (
	"reflect"
	"testing"
)

func TestBuildTreePost(t *testing.T) {
	tests := []struct {
		inorder   []int
		postorder []int
		expected  *TreeNode
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
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
			inorder:   []int{-1},
			postorder: []int{-1},
			expected: &TreeNode{
				Val: -1,
			},
		},
	}

	for _, test := range tests {
		result := buildTreePost(test.inorder, test.postorder)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("buildTreePost(%v, %v) = %v; want %v", test.inorder, test.postorder, result, test.expected)
		}
	}
}
