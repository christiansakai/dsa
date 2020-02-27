package solution

func Solve(m, n int) int {
	rows := n
	cols := m

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 && c == 0 || r == 0 || c == 0 {
				matrix[r][c] = 1
			} else {
				fromTop := matrix[r-1][c]
				fromLeft := matrix[r][c-1]

				matrix[r][c] = fromLeft + fromTop
			}
		}
	}

	return matrix[rows-1][cols-1]
}
