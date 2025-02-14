package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodes := []*TreeNode{root}
	res := [][]int{}

	for len(nodes) != 0 {
		var nInt []int
		var newNodes []*TreeNode

		for _, n := range nodes {

			nInt = append(nInt, n.Val)

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
		nodes = newNodes
	}

	return res
}
