package linked

func NewListNode(arr []int) *ListNode {
	var first *ListNode
	var last *ListNode

	for i, v := range arr {
		nn := &ListNode{Val: v}
		if i == 0 {
			first = nn
		} else {
			last.Next = nn
		}
		last = nn
	}
	return first
}

func MergeTwoLists(first *ListNode, sec *ListNode) *ListNode {
	f := first
	s := sec
	var frn *ListNode
	var lrn *ListNode
	for f != nil || s != nil {
		ln := &ListNode{}

		if f != nil && (s == nil || f.Val < s.Val) {
			ln.Val = f.Val
			f = f.Next
		} else {
			ln.Val = s.Val
			s = s.Next
		}

		if frn == nil {
			frn = ln
		} else {
			lrn.Next = ln
		}
		lrn = ln
	}

	return frn
}
