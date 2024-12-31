package btree

func isSymmetric(root *TreeNode) bool {
	return isInvTree(root.Left, root.Right)
}

func isInvTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isInvTree(p.Left, q.Right) && isInvTree(p.Right, q.Left)
}
