package solution

func Solve(ll *LinkedList) int {
	result := recurse(ll.Head)
	return result - 1 // Adjustment for tail
}

func recurse(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + recurse(node.Next)
}
