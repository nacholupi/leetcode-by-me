package btree

import (
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: true,
		},
		{
			name: "two levels",
			args: args{
				root: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 1},
				},
			},
			want: true,
		},
		{
			name: "two levels with not equals",
			args: args{
				root: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
		{
			name: "two levels with not equals(nil)",
			args: args{
				root: &TreeNode{
					Val:  0,
					Left: &TreeNode{Val: 1},
				},
			},
			want: false,
		},
		{
			name: "three levels",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
