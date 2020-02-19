package solution

func Solve(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	subProb := Solve(head.Next)

	head.Next.Next = head
	head.Next = nil

	return subProb
}
