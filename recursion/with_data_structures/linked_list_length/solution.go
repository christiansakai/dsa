package solution

func Solve(head *Node) int {
	if head == nil {
		return 0
	}

	return 1 + Solve(head.Next)
}
