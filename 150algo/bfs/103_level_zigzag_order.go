package bfs

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodes := []*TreeNode{root}
	res := [][]int{}
	leftR := 1

	for len(nodes) != 0 {
		var nInt []int
		var newNodes []*TreeNode

		for i := range nodes {
			idx := i
			if leftR < 0 {
				idx = len(nodes) - 1 - i
			}

			nInt = append(nInt, nodes[idx].Val)

			n := nodes[i]
			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}

			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}

		if len(nInt) > 0 {
			res = append(res, nInt)
		}

		leftR = -leftR

		nodes = newNodes
	}

	return res
}
