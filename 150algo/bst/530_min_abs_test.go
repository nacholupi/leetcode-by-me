package bst

import (
	"testing"
)

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "test 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1,
		},
		{
			name: "test 2",
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: 1,
		},
		{
			name: "test 3",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1,
		},
		// add this Test [236,104,701,null,227,null,911]
		{
			name: "test 4",
			root: &TreeNode{
				Val: 236,
				Left: &TreeNode{
					Val:  104,
					Left: nil,
					Right: &TreeNode{
						Val:   227,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  701,
					Left: nil,
					Right: &TreeNode{
						Val:   911,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMinimumDifference(tt.root)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
