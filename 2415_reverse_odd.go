package array

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	nodes := 0

	for len(queue) != 0 {

		top := queue[0]
		queue = queue[1:]
		if top.Left == nil {
			break
		}
		nodes++
		queue = append(queue, top.Left, top.Right)

		level := int(logBase2(float64(nodes)))

		if sumPowersOfTwo(level) == nodes && level%2 == 0 {
			for i := 0; i < len(queue)/2; i++ {
				j := len(queue) - 1 - i
				oldLeft := queue[i].Val
				queue[i].Val = queue[j].Val
				queue[j].Val = oldLeft
			}
		}
	}
	return root
}

func logBase2(y float64) float64 {
	return math.Log(y) / math.Log(2)
}

func sumPowersOfTwo(n int) int {
	return int(math.Pow(2, float64(n+1))) - 1
}
