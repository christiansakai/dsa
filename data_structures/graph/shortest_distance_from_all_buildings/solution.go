package solution

import "math"

func Solve(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	var min float64 = math.Inf(1)
	for _, row := range grid {
		for _, col := range row {
			if grid[row][col] == 0 {
				visited := map[int]map[int]bool{}
				distance := getDistance(grid, row, col, visited)
				if distance == -1 {
					return -1
				}

				min = math.Min(min, distance)
			}
		}
	}

	return int(min)
}

func getDistance(grid [][]int, row, col int, visited map[int]map[int]bool) (int, int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return 0, 0
	}

	if grid[row][col] == 2 {
		if _, ok := visited[row]; !ok {
			visited[row] = map[int]bool{}
		}

		visited[row][col] = true

		return 0, 0
	}

	if grid[row][col] == 1 {
		if _, ok := visited[row]; !ok {
			visited[row] = map[int]bool{}
		}

		visited[row][col] = true

		return 1, 1
	}

	houseFound, distance := getDistance(grid, row-1, col)
	min = math.Min(min, fromTop)

	fromRight := getDistance(grid, row, col+1)
	min = math.Min(min, fromRight)

	fromBottom := getDistance(grid, row+1, col)
	min = math.Min(min, fromTop)

	fromLeft := getDistance(grid, row, col-1)
	min = math.Min(min, fromTop)
}
