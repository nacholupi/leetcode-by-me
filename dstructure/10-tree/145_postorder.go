package tree

/* 94. Binary Tree Inorder Traversal*/
func postorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	traversalP(root, &r)
	return r
}

func traversalP(root *TreeNode, r *[]int) {
	if root == nil {
		return
	}

	traversalP(root.Left, r)
	traversalP(root.Right, r)
	*r = append(*r, root.Val)
}

func postorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stackNode := stack{root}
	for !stackNode.isEmpty() {
		r := stackNode.peek()
		if r.Left != nil {
			stackNode.push(r.Left)
			r.Left = nil
			continue
		}

		if r.Right != nil {
			stackNode.push(r.Right)
			r.Right = nil
			continue
		}

		res = append(res, stackNode.pop().Val)
	}
	return res
}
