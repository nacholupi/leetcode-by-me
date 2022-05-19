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

type stack []*TreeNode

func (s *stack) push(r *TreeNode) {
	*s = append(*s, r)
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) peek() *TreeNode {
	i := len(*s) - 1
	e := (*s)[i]
	return e
}

func (s *stack) pop() *TreeNode {
	i := len(*s) - 1
	e := (*s)[i]
	*s = (*s)[:i]
	return e
}

func inorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stackNode := stack{root}
	for !stackNode.isEmpty() {
		r := stackNode.peek()
		if r.Left == nil {
			res = append(res, stackNode.pop().Val)
			if r.Right != nil {
				stackNode.push(r.Right)
			}
		}
		if r.Left != nil {
			stackNode.push(r.Left)
			r.Left = nil
		}
	}
	return res
}
