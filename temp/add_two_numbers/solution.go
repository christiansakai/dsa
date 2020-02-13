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

	sum := l1.Val + l2.Val
	if sum <= 9 {
		return &ListNode{
			Val:  sum,
			Next: Solve(l1.Next, l2.Next),
		}
	}

	remainder := sum - 10
	if l1.Next == nil {
		l1.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	} else {
		l1.Next.Val += 1
		if l1.Next.Val > 9 {
			normalize(l1.Next)
		}
	}

	return &ListNode{
		Val:  remainder,
		Next: Solve(l1.Next, l2.Next),
	}
}

func normalize(l *ListNode) {
	if l == nil {
		return
	}

	remainder := l.Val - 10
	l.Val = remainder

	if l.Next == nil {
		l.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	} else {
		l.Next.Val = 1
	}
}
