package tree

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "easy case",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 2},
			},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			name: "easy case - right",
			args: args{
				root: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 1},
				},
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 2},
			},
			want: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name: "root in the middle",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 3},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "subtree",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{Val: 5},
				},
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 2},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
