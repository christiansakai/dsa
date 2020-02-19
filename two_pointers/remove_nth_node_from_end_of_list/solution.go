package solution

func Solve(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	if head.Next == nil && n == 1 {
		return nil
	}

	recurse(head, head.Next, n)

	return head
}

func recurse(curr, next *ListNode, n int) int {
	if next.Next == nil {
		if n == 1 {
			curr.Next = nil
			return -1 // signifies not to do anything
		} else if n == 2 {
			curr.Val = next.Val
			curr.Next = nil
		}

		return 1
	}

	subProb := recurse(curr.Next, next.Next, n)
	if subProb == -1 {
		return -1
	}

	currIndex := 1 + subProb
	if currIndex != n {
		return currIndex
	}

	curr.Next = next.Next
	return currIndex
}
