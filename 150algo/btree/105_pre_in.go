package btree

func buildTree(preorder []int, inorder []int) *TreeNode {

	inMapIdx := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inMapIdx[v] = i
	}

	i := 0

	var tree func(lIxd int, rIdx int) *TreeNode

	tree = func(lIxd int, rIdx int) *TreeNode {
		if lIxd > rIdx {
			return nil
		}

		cur := &TreeNode{Val: preorder[i]}
		i++
		idx := inMapIdx[cur.Val]
		cur.Left = tree(lIxd, idx-1)
		cur.Right = tree(idx+1, rIdx)

		return cur
	}

	return tree(0, len(inorder)-1)
}
