package reverselist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}
