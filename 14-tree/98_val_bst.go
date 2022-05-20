package tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, math.MinInt, math.MaxInt)
}

func isValidNode(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}
	validLeft := node.Left == nil || (node.Left.Val < node.Val && node.Left.Val > minVal)
	validRight := node.Right == nil || (node.Right.Val > node.Val && node.Right.Val < maxVal)
	if !(validLeft && validRight) {
		return false
	}

	return isValidNode(node.Left, minVal, node.Val) && isValidNode(node.Right, node.Val, maxVal)
}
