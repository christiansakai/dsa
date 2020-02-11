package solution

func Solve(ll *LinkedList) *LinkedList {
	if ll.Length <= 1 {
		return ll
	}

	// reversed := recurse(ll.Head.Next)
	// ll.Head.Next = reversed

	return ll
}

func recurse(node *Node) *Node {
	if node.Value == nil { // Encounter dummy tail
		return node
	}

	subProb := recurse(node.Next)
	subProb.Next = node

	node.Prev = node
	return subProb
}
