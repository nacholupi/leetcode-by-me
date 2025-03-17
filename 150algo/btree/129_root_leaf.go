package btree

func sumNumbers(root *TreeNode) int {

	count := 0

	var dfs func(root *TreeNode, parentVal int)
	dfs = func(root *TreeNode, parentVal int) {

		val := parentVal*10 + root.Val
		if root.Left == nil && root.Right == nil {
			count += val
		}

		if root.Left != nil {
			dfs(root.Left, val)
		}

		if root.Right != nil {
			dfs(root.Right, val)
		}
	}

	dfs(root, 0)
	return count
}
