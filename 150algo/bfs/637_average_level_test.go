package bfs

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected []float64
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: []float64{3, 14.5, 11},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: []float64{1, 2.5, 5},
		},
		{
			root:     nil,
			expected: []float64{},
		},
	}

	for _, test := range tests {
		result := averageOfLevels(test.root)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, result)
		}
	}
}
