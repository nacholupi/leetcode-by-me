package btree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	newTargetSum := targetSum - root.Val

	if isLeaf(root) && newTargetSum == 0 {
		return true
	}

	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
