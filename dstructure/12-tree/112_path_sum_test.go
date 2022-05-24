package tree

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "zero case",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "one case",
			args: args{
				root:      &TreeNode{Val: 1},
				targetSum: 1,
			},
			want: true,
		},
		{
			name: "two level case, right path",
			args: args{
				root:      &TreeNode{Val: 1},
				targetSum: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
