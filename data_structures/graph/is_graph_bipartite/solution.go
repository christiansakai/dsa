package solution

func SolveIterative(graph [][]int) bool {
	// Also double as a visited map
	colors := make([]int, len(graph))
	for i := 0; i < len(colors); i++ {
		colors[i] = -1 // not colored yet
	}

	for i := 0; i < len(colors); i++ {
		if colors[i] == -1 {
			// Iterative DFS coloring
			stack := []int{i}
			colors[i] = 0

			for len(stack) > 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				neighbors := graph[node]
				for _, neighbor := range neighbors {
					// if neighbor is not colored yet
					if colors[neighbor] == -1 {
						stack = append(stack, neighbor)

						// color all neighbor differently
						if colors[node] == 0 {
							colors[neighbor] = 1
						} else {
							colors[neighbor] = 0
						}
					} else {
						// if neighbor is already colored
						if colors[neighbor] == colors[node] {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

func SolveRecursive(graph [][]int) bool {
	colors := make([]int, len(graph))
	for i := 0; i < len(colors); i++ {
		colors[i] = -1 // not colored yet
	}

	for i := 0; i < len(colors); i++ {
		if colors[i] == -1 && !colorGraph(i, 0, &colors, graph) {
			return false
		}
	}

	return true
}

func colorGraph(node, color int, colors *[]int, graph [][]int) bool {
	if (*colors)[node] != -1 {
		if (*colors)[node] == color {
			return true
		}

		return false
	}

	(*colors)[node] = color

	neighbors := graph[node]
	for _, n := range neighbors {
		if color == 0 {
			if !colorGraph(n, 1, colors, graph) {
				return false
			}
		} else {
			if !colorGraph(n, 0, colors, graph) {
				return false
			}
		}
	}

	return true
}
