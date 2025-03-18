package btree

import (
	"testing"
)

func TestBSTIterator(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	it := Constructor(root)

	expected := []int{3, 7, 9, 15, 20}
	for _, val := range expected {
		if !it.HasNext() {
			t.Fatalf("expected HasNext() to be true, but got false")
		}
		if next := it.Next(); next != val {
			t.Fatalf("expected Next() to be %d, but got %d", val, next)
		}
	}

	if it.HasNext() {
		t.Fatalf("expected HasNext() to be false, but got true")
	}
}
