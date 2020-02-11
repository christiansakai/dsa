package solution

type Graph struct {
	Dict map[int][]int
}

func New() *Graph {
	dict := make(map[int][]int)

	return &Graph{dict}
}

func (g *Graph) AddEdge(u, v int) {
	if _, ok := g.Dict[u]; !ok {
		g.Dict[u] = []int{}
	}

	g.Dict[u] = append(g.Dict[u], v)
}
