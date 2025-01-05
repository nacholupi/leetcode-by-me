package linked

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	lrn := dummy

	for list1 != nil || list2 != nil {
		ln := &ListNode{}

		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			ln.Val = list1.Val
			list1 = list1.Next
		} else {
			ln.Val = list2.Val
			list2 = list2.Next
		}

		lrn.Next = ln
		lrn = ln
	}

	return dummy.Next
}
