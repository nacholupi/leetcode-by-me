package tree

// TODO: lousy
func isSymmetric(root *TreeNode) bool {
	return symChilds([]*TreeNode{root})
}

func symChilds(nodes []*TreeNode) bool {
	childs := make([]*TreeNode, len(nodes)*2)
	withChilds := false

	if len(nodes) == 1 {
		if nodes[0].Left != nil {
			childs[0] = nodes[0].Left
			withChilds = true
		}

		if nodes[0].Right != nil {
			childs[1] = nodes[0].Right
			withChilds = true
		}

		if !equals(childs[0], childs[1]) {
			return false
		}
	}

	for i := 0; i < len(nodes); i++ {

		n := len(nodes) - 1 - i
		if nodes[i] != nil && nodes[i].Left != nil {
			childs[i*2] = nodes[i].Left
			withChilds = true
		}

		if nodes[n] != nil && nodes[n].Right != nil {
			childs[n*2+1] = nodes[n].Right
			withChilds = true
		}

		if !equals(childs[i*2], childs[n*2+1]) {
			return false
		}

		if nodes[i] != nil && nodes[i].Right != nil {
			childs[i*2+1] = nodes[i].Right
			withChilds = true
		}

		if nodes[n] != nil && nodes[n].Left != nil {
			childs[n*2] = nodes[n].Left
			withChilds = true
		}

		if !equals(childs[i*2+1], childs[n*2]) {
			return false
		}
	}

	if !withChilds {
		return true
	}

	return symChilds(childs)
}

func equals(a, b *TreeNode) bool {
	if (a == nil && b == nil) || (a != nil && b != nil && a.Val == b.Val) {
		return true
	}
	return false
}
