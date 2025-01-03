package linked

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := false
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil || l2 != nil || carry {

		nl1, nl2, nt := 0, 0, 0

		if l1 != nil {
			nl1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			nl2 = l2.Val
			l2 = l2.Next
		}

		if carry {
			nt = 1
			carry = false
		}

		val := nl1 + nl2 + nt
		if val > 9 {
			carry = true
		}

		l := &ListNode{Val: val % 10}
		tail.Next = l
		tail = l
	}
	return dummy.Next
}
