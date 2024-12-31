package btree

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftD := 1
	l := root.Left
	for l != nil {
		leftD++
		l = l.Left
	}

	rightD := 1
	r := root.Right
	for r != nil {
		rightD++
		r = r.Right
	}

	if leftD == rightD {
		return int(math.Pow(2, float64(leftD))) - 1
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}
