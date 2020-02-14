package solution

func Solve(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	if l1 == nil && l2 == nil {
		return head
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	carry, remainder := add(l1.Val, l2.Val)

	head = &ListNode{
		Val:  remainder,
		Next: nil,
	}

	var temp *ListNode = head
	l1 = l1.Next
	l2 = l2.Next

	for l1 != nil || l2 != nil {
		if l1 == nil {
			tempCarry, remainder := add(l2.Val, carry)
			carry = tempCarry

			newNode := &ListNode{
				Val:  remainder,
				Next: nil,
			}

			temp.Next = newNode

			temp = temp.Next
			l2 = l2.Next
		} else if l2 == nil {
			tempCarry, remainder := add(l1.Val, carry)
			carry = tempCarry

			newNode := &ListNode{
				Val:  remainder,
				Next: nil,
			}

			temp.Next = newNode

			temp = temp.Next
			l1 = l1.Next
		} else {
			tempCarry, remainder := add(carry+l1.Val, l2.Val)
			carry = tempCarry

			newNode := &ListNode{
				Val:  remainder,
				Next: nil,
			}

			temp.Next = newNode

			temp = temp.Next
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	if carry > 0 {
		newNode := &ListNode{
			Val:  carry,
			Next: nil,
		}

		temp.Next = newNode
		temp = temp.Next
	}

	return head
}

func add(n1, n2 int) (int, int) {
	sum := n1 + n2
	if sum >= 10 {
		return 1, sum - 10
	}

	return 0, sum
}
