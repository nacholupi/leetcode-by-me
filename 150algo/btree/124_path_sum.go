package btree

import (
	"math"
)

func maxPathSum(root *TreeNode) int {

	if root == nil {
		return math.MinInt32
	}

	maxlVal := maxPathSum(root.Left)
	maxrVal := maxPathSum(root.Right)

	l := math.MinInt32
	if root.Left != nil {
		l = root.Left.Val
	}

	r := math.MinInt32
	if root.Right != nil {
		r = root.Right.Val
	}

	sumAll := root.Val + l + r
	root.Val += max(0, max(l, r))
	return max(max(sumAll, root.Val), max(maxlVal, maxrVal))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
