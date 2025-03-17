package btree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		root     *Node
		expected *Node
	}{
		{
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},
			},
			expected: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
						Next: &Node{
							Val: 5,
						},
					},
					Right: &Node{
						Val: 5,
						Next: &Node{
							Val: 7,
						},
					},
					Next: &Node{
						Val: 3,
					},
				},
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := connect(test.root)
		fmt.Println(result)
		assert.Equal(t, fmt.Sprintf("%v", test.expected), fmt.Sprintf("%v", result))
	}
}
