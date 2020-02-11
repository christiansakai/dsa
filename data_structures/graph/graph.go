package graph

type Graph struct {
	Dict map[interface{}][]interface{}
}

func New() *Graph {
	dict := make(map[interface{}][]interface{})

	return &Graph{dict}
}

func (g *Graph) AddEdge(u, v interface{}) {
	if _, ok := g.Dict[u]; !ok {
		g.Dict[u] = []interface{}{}
	}

	g.Dict[u] = append(g.Dict[u], v)
}
