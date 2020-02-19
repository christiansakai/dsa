package solution

func Solve(head *ListNode) bool {
	middle := head
	end := head

	for {
		if end == nil || end.Next == nil {
			break
		}

		middle = middle.Next
		end = end.Next.Next
	}

}
