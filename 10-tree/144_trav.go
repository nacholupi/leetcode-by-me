package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return recPreorderTraversal(root)
}

func recPreorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	addVal(root, res)
	return *res
}

func addVal(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	addVal(root.Left, result)
	addVal(root.Right, result)
}
