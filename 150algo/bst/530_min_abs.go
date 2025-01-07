package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	queue := []*TreeNode{root}
	res := 100000

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.Left != nil {
			i := n.Left
			val := 0
			for i != nil {
				val = i.Val
				i = i.Right
			}
			if n.Val-val < res {
				res = n.Val - val
			}

			queue = append(queue, n.Left)
		}

		if n.Right != nil {
			i := n.Right
			val := 0
			for i != nil {
				val = i.Val
				i = i.Left
			}
			if val-n.Val < res {
				res = val - n.Val
			}

			queue = append(queue, n.Right)
		}
	}
	return res
}

/*
func getMinimumDifference(root *TreeNode) int {
    arr := dfs(root, []int{})
    minim := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        minim = min(minim, arr[i] - arr[i-1])
    }
    return minim
}


func dfs(root *TreeNode, arr []int) []int {
    if root == nil {
        return arr
    }

    arr = dfs(root.Left, arr)
    arr = append(arr, root.Val)
    arr = dfs(root.Right, arr)
    return arr
}
*/
