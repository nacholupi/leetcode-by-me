package btree

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "example 2",
			root: &TreeNode{
				Val: 0,
			},
			want: &TreeNode{
				Val: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			if !reflect.DeepEqual(tt.root, tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.root, tt.want)
			}
		})
	}
}
