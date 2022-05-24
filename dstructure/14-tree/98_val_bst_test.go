package tree

import "testing"

func Test_isValidBST(t *testing.T) {
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
				root: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "two levels",
			args: args{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2}},
			},
			want: true,
		},
		{
			name: "3 levels",
			args: args{
				root: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: false,
		},
		/* [5,4,6,null,null,3,7] */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
