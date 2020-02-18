package solution

func Solve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextNext := head.Next.Next
	next := head.Next

	head.Next = Solve(nextNext)
	next.Next = head

	return next
}
