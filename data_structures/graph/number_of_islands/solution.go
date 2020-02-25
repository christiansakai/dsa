package solution

func Solve(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				islands += 1

				recurse(grid, r, c)
			}
		}
	}

	return islands
}

func recurse(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	if grid[row][col] == '1' {
		grid[row][col] = '0'
	}

	neighbors := getNeighbors(grid, row, col)
	for _, n := range neighbors {
		recurse(grid, n.row, n.col)
	}
}

type coordinate struct {
	row int
	col int
}

func getNeighbors(grid [][]byte, row, col int) []coordinate {
	neighbors := []coordinate{}

	// Get top bottom
	for i := -1; i <= 1; i++ {
		if i != 0 {
			neighborRow := row + i
			neighborCol := col

			if neighborRow >= 0 && neighborRow < len(grid) && neighborCol >= 0 && neighborCol < len(grid[0]) {
				neighbors = append(neighbors, coordinate{
					row: neighborRow,
					col: neighborCol,
				})
			}
		}
	}

	// Get top bottom
	for i := -1; i <= 1; i++ {
		if i != 0 {
			neighborRow := row
			neighborCol := col + i

			if neighborRow >= 0 && neighborRow < len(grid) && neighborCol >= 0 && neighborCol < len(grid[0]) {
				neighbors = append(neighbors, coordinate{
					row: neighborRow,
					col: neighborCol,
				})
			}
		}
	}

	return neighbors
}
