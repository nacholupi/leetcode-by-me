package bst

func kthSmallest(root *TreeNode, k int) int {
	result := 0
	count := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorder(node.Right)
	}

	inorder(root)
	return result
}
