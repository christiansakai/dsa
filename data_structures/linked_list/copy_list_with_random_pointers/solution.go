package solution

func Solve(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	visited := map[*ListNode]*ListNode{}
	return recurse(head, visited)
}

func recurse(node *ListNode, visited map[*ListNode]*ListNode) *ListNode {
	if node == nil {
		return nil
	}

	if cloned, ok := visited[node]; ok {
		return cloned
	}

	cloned := &ListNode{Val: node.Val}
	visited[node] = cloned

	cloned.Next = recurse(node.Next, visited)
	cloned.Random = recurse(node.Random, visited)

	return cloned
}
