package solution

func solve(graph [][]int) int {
	if len(graph) == 0 {
		return 0
	}

	visited := map[int]bool{}
	circleOfFriends := 0

	for f := 0; f < len(graph); f++ {
		if !visited[f] {
			visit(graph, f, visited)
			circleOfFriends += 1
		}
	}

	return circleOfFriends
}

func visit(graph [][]int, p int, visited map[int]bool) {
	if visited[p] {
		return
	}

	visited[p] = true

	for _, f := range getFriends(graph, p) {
		if !visited[f] {
			visit(graph, f, visited)
		}
	}
}

func getFriends(graph [][]int, p int) []int {
	friends := []int{}

	for f := 0; f < len(graph[p]); f++ {
		if graph[p][f] == 1 && p != f {
			friends = append(friends, f)
		}
	}

	return friends
}
