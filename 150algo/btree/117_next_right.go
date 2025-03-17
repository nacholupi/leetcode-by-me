package btree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("Val %d, Next %d, Left %v, Right %v", n.Val, n.Next.Val, n.Left, n.Right)
}

func connect(root *Node) *Node {

	lastNodeLevel := make(map[int]*Node)

	var bfs func(root *Node, level int)

	bfs = func(root *Node, level int) {
		if root == nil {
			return
		}
		root.Next = lastNodeLevel[level]
		lastNodeLevel[level] = root
		bfs(root.Right, level+1)
		bfs(root.Left, level+1)
	}

	bfs(root, 0)
	return root
}
