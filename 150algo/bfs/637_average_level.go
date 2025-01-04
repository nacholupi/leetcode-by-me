package bfs

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	nodes := []*TreeNode{root}
	res := []float64{}

	for len(nodes) != 0 {
		var newNodes []*TreeNode
		partialSum := 0
		partialCount := 0

		for _, n := range nodes {
			partialSum += n.Val
			partialCount++

			if n.Left != nil {
				newNodes = append(newNodes, n.Left)
			}

			if n.Right != nil {
				newNodes = append(newNodes, n.Right)
			}
		}

		if partialCount > 0 {
			res = append(res, float64(partialSum)/float64(partialCount))
		}
		nodes = newNodes
	}

	return res
}
