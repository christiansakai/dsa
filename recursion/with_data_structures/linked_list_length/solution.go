package solution

import dll "dsa/data_structures/doubly_linked_list"

func Solve(ll *dll.LinkedList) int {
	result := recurse(ll.Head)
	return result - 1 // Adjustment for tail
}

func recurse(node *dll.Node) int {
	if node == nil {
		return 0
	}

	return 1 + recurse(node.Next)
}
