package linked

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
