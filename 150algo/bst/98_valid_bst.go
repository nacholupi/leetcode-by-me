package bst

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	intMin := math.MinInt
	return isValid(root, &intMin)
}

func isValid(root *TreeNode, prev *int) bool {
	if root == nil {
		return true
	}

	if !isValid(root.Left, prev) {
		return false
	}

	if root.Val <= *prev {
		return false
	}

	*prev = root.Val

	if !isValid(root.Right, prev) {
		return false
	}

	return true
}
