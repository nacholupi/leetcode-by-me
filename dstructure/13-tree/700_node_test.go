package tree

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "zero",
			args: args{},
			want: nil,
		},
		{
			name: "one node found",
			args: args{
				root: &TreeNode{Val: 1},
				val:  1,
			},
			want: &TreeNode{Val: 1},
		},
		{
			name: "one node not found",
			args: args{
				root: &TreeNode{Val: 1},
				val:  2,
			},
			want: nil,
		},
		{
			name: "left subtree found",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 2,
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "subtree not found",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
