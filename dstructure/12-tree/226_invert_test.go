package tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
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
			name: "one",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: &TreeNode{Val: 1},
		},
		{
			name: "two levels",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 2},
				},
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "two levels",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{Val: 2},
				},
			},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
