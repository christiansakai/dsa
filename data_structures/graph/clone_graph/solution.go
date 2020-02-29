package solution

func Solve(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	visited := map[*GraphNode]*GraphNode{}
	return recurse(node, visited)
}

func recurse(node *GraphNode, visited map[*GraphNode]*GraphNode) *GraphNode {
	if node == nil {
		return nil
	}

	if cloned, ok := visited[node]; ok {
		return cloned
	}

	cloned := &GraphNode{Val: node.Val}
	visited[node] = cloned

	for _, neighbor := range node.Neighbors {
		clonedNeighbor := recurse(neighbor, visited)
		cloned.Neighbors = append(cloned.Neighbors, clonedNeighbor)
	}

	return cloned
}
