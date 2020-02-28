package solution

func Solve(head *ListNode) {
	temp := head

	for temp != nil {
		temp.Next = reverse(temp.Next)
		temp = temp.Next
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		tempNext := curr.Next

		curr.Next = prev

		prev = curr
		curr = tempNext
	}

	return prev
}
