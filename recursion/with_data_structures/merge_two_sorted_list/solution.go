package solution

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var merged *ListNode = nil
	if l1.Val < l2.Val {
		merged = l1
		merged.Next = Solve(l1.Next, l2)
	} else {
		merged = l2
		merged.Next = Solve(l1, l2.Next)
	}

	return merged
}
