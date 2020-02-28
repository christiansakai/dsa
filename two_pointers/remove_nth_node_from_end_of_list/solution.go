package solution

func Solve(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	// Init fast
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// Init slow
	var prevSlow *ListNode
	slow := head

	// Traverse
	for fast.Next != nil {
		if prevSlow == nil {
			prevSlow = head
		} else {
			prevSlow = prevSlow.Next
		}

		slow = slow.Next
		fast = fast.Next
	}

	// Remove slow
	prevSlow.Next = slow.Next
	slow.Next = nil

	return head
}
