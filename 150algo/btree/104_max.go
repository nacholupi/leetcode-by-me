package btree

func maxDepth(root *TreeNode) int {
	i := 0
	if root == nil {
		return i
	}

	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		i++
		var newNodes []*TreeNode

		for _, n := range nodes {
			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}

			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}

		nodes = newNodes
	}

	return i

}
