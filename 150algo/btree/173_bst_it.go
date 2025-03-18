package btree

type BSTIterator struct {
	nextNode *TreeNode
	stack    []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{nextNode: root}
	it.goLeft()
	return it
}

func (it *BSTIterator) goLeft() {
	for it.nextNode != nil && it.nextNode.Left != nil {
		it.stack = append(it.stack, it.nextNode)
		it.nextNode = it.nextNode.Left
	}
}

func (it *BSTIterator) advance() {
	if it.nextNode != nil && it.nextNode.Right != nil {
		it.nextNode = it.nextNode.Right
		it.goLeft()
	} else if len(it.stack) != 0 {
		it.nextNode = it.stack[len(it.stack)-1]
		it.stack = it.stack[:len(it.stack)-1]
	} else {
		it.nextNode = nil
	}
}

func (it *BSTIterator) Next() int {
	next := it.nextNode.Val
	it.advance()
	return next
}

func (it *BSTIterator) HasNext() bool {
	return it.nextNode != nil
}
