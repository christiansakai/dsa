package solution

func Solve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := Solve(head.Next)
	newHead.Next = head

	return newHead
}
