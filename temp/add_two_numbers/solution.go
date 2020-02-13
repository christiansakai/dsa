package solution

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			sum := l2.Val + carry
			if sum >= 10 {
				carry = 1
			}

			remainder := sum - 10

			newNode := &ListNode{
				Val:  remainder,
				Next: head,
			}

			head = newNode
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			sum := l1.Val + carry
			if sum >= 10 {
				carry = 1
			}

			remainder := sum - 10

			newNode := &ListNode{
				Val:  remainder,
				Next: head,
			}

			head = newNode
			l1 = l1.Next
			continue
		}

		sum := l1.Val + l2.Val + carry
		if sum >= 10 {
			carry = 1
		}

		remainder := sum - 10

		newNode := &ListNode{
			Val:  remainder,
			Next: head,
		}

		head = newNode
		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		newNode := &ListNode{
			Val:  carry,
			Next: head,
		}

		head = newNode
	}

	return head
}
