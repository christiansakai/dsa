package solution

func Solve(node *ListNode) {
	if node.Next.Next == nil {
		node.Val = node.Next.Val
		node.Next = nil
		return
	}

	node.Val = node.Next.Val
	Solve(node.Next)
}
