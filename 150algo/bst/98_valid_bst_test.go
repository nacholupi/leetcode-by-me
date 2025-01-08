package bst

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	var tests = []struct {
		desc  string
		input *TreeNode

		expected bool
	}{
		{
			desc: "test1",
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			desc: "test2",
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			desc: "test3",
			input: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: false,
		},
		// next treenode to test [32,26,47,19,null,null,56,null,27]
		{
			desc: "test4",
			input: &TreeNode{
				Val: 32,
				Left: &TreeNode{
					Val: 26,
					Left: &TreeNode{
						Val: 19,
						Right: &TreeNode{
							Val: 27,
						},
					},
				},
				Right: &TreeNode{
					Val: 47,
					Right: &TreeNode{
						Val: 56,
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		if output := isValidBST(test.input); output != test.expected {
			t.Errorf("Test (%s) Failed: %v expected, received: %v", test.desc, test.expected, output)
		}
	}
}
