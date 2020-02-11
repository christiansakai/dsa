package solution

import dll "dsa/data_structures/doubly_linked_list"

func Solve(ll *dll.LinkedList) *dll.LinkedList {
	if ll.Length == 0 {
		return ll
	}

	reversed := recurse(ll.Head.Next)
	ll.Head.Next = reversed

	return ll
}

func recurse(node *dll.Node) *dll.Node {
	if node.Value == nil { // Encounter dummy tail
		return node
	}

	subProb := recurse(node.Next)
	subProb.Next = node

	node.Prev = node
	return subProb
}
