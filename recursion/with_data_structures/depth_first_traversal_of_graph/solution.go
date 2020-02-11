package solution

func Solve(g *Graph) []int {
	collect := []int{}
	visited := map[int]bool{}

	recurse(g, 0, visited, &collect)

	return collect
}

func recurse(g *Graph, node int, visited map[int]bool, collect *[]int) {
	if visited[node] {
		return
	}

	*collect = append(*collect, node)
	visited[node] = true

	for _, neighbor := range g.Dict[node] {
		recurse(g, neighbor, visited, collect)
	}
}
