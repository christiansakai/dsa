package solution

import "dsa/data_structures/graph"

func Solve(g *graph.Graph) []int {
	collect := []int{}
	visited := map[int]bool{}

	recurse(g, 0, visited, &collect)

	return collect
}

func recurse(g *graph.Graph, node int, visited map[int]bool, collect *[]int) {
	if visited[node] {
		return
	}

	*collect = append(*collect, node)
	visited[node] = true

	for _, neighbor := range g.Dict[node] {
		neighborInt, _ := neighbor.(int)
		recurse(g, neighborInt, visited, collect)
	}
}
