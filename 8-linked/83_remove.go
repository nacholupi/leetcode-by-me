package linked

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
