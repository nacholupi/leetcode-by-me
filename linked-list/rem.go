package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		head = head.Next
	}

	h := head
	for h != nil {
		if h.Next != nil && h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	for h != nil {
		if h.Next != nil && h.Val == h.Next.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}

	return head
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	if head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head.Next
	fast := head.Next.Next

	for {

		if slow == fast {
			return true
		}

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}
