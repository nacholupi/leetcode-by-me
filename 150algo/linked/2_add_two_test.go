package linked

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			l1:       createList([]int{2, 4, 3}),
			l2:       createList([]int{5, 6, 4}),
			expected: createList([]int{7, 0, 8}),
		},
		{
			l1:       createList([]int{0}),
			l2:       createList([]int{0}),
			expected: createList([]int{0}),
		},
		{
			l1:       createList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       createList([]int{9, 9, 9, 9}),
			expected: createList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		result := addTwoNumbers(test.l1, test.l2)
		if !compareLists(result, test.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", listToSlice(test.l1), listToSlice(test.l2), listToSlice(result), listToSlice(test.expected))
		}
	}
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func listToSlice(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}
