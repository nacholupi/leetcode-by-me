package bfs

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	nodes := []*TreeNode{root}
	res := []int{}

	for len(nodes) != 0 {
		var newNodes []*TreeNode

		for _, n := range nodes {

			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}

			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}

		res = append(res, nodes[len(nodes)-1].Val)
		nodes = newNodes
	}

	return res
}
