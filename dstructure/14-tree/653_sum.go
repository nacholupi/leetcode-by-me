package tree

/* 653. Two Sum IV - Input is a BST */
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return dfs(root, &m, k)
}

func dfs(root *TreeNode, m *map[int]struct{}, k int) bool {
	if root == nil {
		return false
	}

	if _, ok := (*m)[k-root.Val]; ok {
		return true
	}
	(*m)[root.Val] = struct{}{}

	return dfs(root.Left, m, k) || dfs(root.Right, m, k)
}
