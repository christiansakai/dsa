package solution

func Solve(g *Graph) []int {
	nodes := collectAllNodes(g)
	result := []int{}
	visited := map[int]bool{}
	recurse(g, nodes, visited, &result)

	return result
}

func collectAllNodes(g *Graph) map[int]bool {
	nodes := map[int]bool{}
	for node, neighbors := range g.Dict {
		nodes[node] = true

		for _, neighbor := range neighbors {
			nodes[neighbor] = true
		}
	}

	return nodes
}

func recurse(g *Graph, nodes map[int]bool, visited map[int]bool, result *[]int) {
	node := findUnvisitedNodeWithZeroInDegree(g, visited, nodes)
	if node == -1 {
		return
	}

	*result = append(*result, node)
	visited[node] = true
	removeNodeFromOtherInDegrees(g, node)

	recurse(g, nodes, visited, result)
}

func findUnvisitedNodeWithZeroInDegree(g *Graph, visited map[int]bool, nodes map[int]bool) int {
	inDegrees := map[int]int{}

	for _, neighbors := range g.Dict {
		for _, neighbor := range neighbors {
			inDegrees[neighbor] = 1
		}
	}

	for node, _ := range nodes {
		if inDegrees[node] == 0 && !visited[node] {
			return node
		}
	}

	return -1
}

func removeNodeFromOtherInDegrees(g *Graph, node int) {
	delete(g.Dict, node)
}
