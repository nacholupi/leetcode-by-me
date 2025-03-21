package btree

/* best solution
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != r {
		if l != nil && r != nil {
			return root
		} else if l != nil {
			return l
		} else {
			return r
		}
	} else {
		return r
	}
}
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var firstPath []*TreeNode
	var secPath []*TreeNode

	var findPath func(root *TreeNode, path []*TreeNode)
	findPath = func(root *TreeNode, path []*TreeNode) {
		if root == nil {
			return
		}

		path = append(path, root)

		if root == p || root == q {
			if len(firstPath) == 0 {
				firstPath = append([]*TreeNode(nil), path...)
			} else {
				secPath = append([]*TreeNode(nil), path...)
			}
		}

		if len(secPath) > 0 {
			return
		}
		findPath(root.Left, path)

		if len(secPath) > 0 {
			return
		}
		findPath(root.Right, path)
	}

	findPath(root, []*TreeNode{})

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var common *TreeNode
	for i := 0; i < min(len(firstPath), len(secPath)); i++ {
		if firstPath[i] != secPath[i] {
			break
		}
		common = firstPath[i]
	}

	return common
}
