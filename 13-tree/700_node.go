package tree

/* 700. Search in a Binary Search Tree */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Left != nil && val < root.Val {
		return searchBST(root.Left, val)
	}
	if root.Right != nil && val > root.Val {
		return searchBST(root.Right, val)
	}

	return nil
}
