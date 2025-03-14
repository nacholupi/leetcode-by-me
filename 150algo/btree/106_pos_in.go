package btree

func buildTreePost(inorder []int, postorder []int) *TreeNode {

	inMapIdx := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inMapIdx[v] = i
	}

	i := len(postorder) - 1

	var tree func(lIxd int, rIdx int) *TreeNode

	tree = func(lIxd int, rIdx int) *TreeNode {
		if lIxd > rIdx {
			return nil
		}

		cur := &TreeNode{Val: postorder[i]}
		i--
		idx := inMapIdx[cur.Val]
		cur.Right = tree(idx+1, rIdx)
		cur.Left = tree(lIxd, idx-1)

		return cur
	}

	return tree(0, len(inorder)-1)
}
