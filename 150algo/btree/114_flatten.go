package btree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	node := root

	for node != nil {
		rSide := node.Right
		node.Right = node.Left
		node.Left = nil

		rightMostChild := node
		for rightMostChild.Right != nil {
			rightMostChild = rightMostChild.Right
		}

		rightMostChild.Right = rSide
		node = node.Right
	}
}
