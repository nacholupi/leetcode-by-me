package tree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "zero",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "one",
			args: args{
				root: &TreeNode{
					Val:   10,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{10},
		},
		{
			name: "left with val",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			want: []int{10, 8},
		},
		{
			name: "left with val",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   9,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:  11,
							Left: nil,
							Right: &TreeNode{
								Val:   12,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: nil,
				},
			},
			want: []int{10, 5, 9, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
