package tree

/* 94. Binary Tree Inorder Traversal*/
func inorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	traversal(root, &r)
	return r
}

func traversal(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, r)
	*r = append(*r, root.Val)
	traversal(root.Right, r)
}
