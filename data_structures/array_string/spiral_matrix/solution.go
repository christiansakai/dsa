package solution

func Solve(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	result := []int{}

	element := len(matrix) * len(matrix[0])

	l := 0
	r := len(matrix[0]) - 1
	u := 0
	d := len(matrix) - 1

	direction := "lr"

	i := 0
	for i < element {
		if direction == "lr" {
			runLR(matrix, u, l, r, &result)

			u += 1
			direction = "ud"
			i += r - l + 1
		} else if direction == "ud" {
			runUD(matrix, r, u, d, &result)

			r -= 1
			direction = "rl"
			i += d - u + 1
		} else if direction == "rl" {
			runRL(matrix, d, r, l, &result)

			d -= 1
			direction = "du"
			i += r - l + 1
		} else {
			runDU(matrix, l, d, u, &result)

			l += 1
			direction = "lr"
			i += d - u + 1
		}
	}

	return result
}

func runLR(matrix [][]int, row, from, to int, result *[]int) {
	for i := from; i <= to; i++ {
		*result = append(*result, matrix[row][i])
	}
}

func runUD(matrix [][]int, col, from, to int, result *[]int) {
	for i := from; i <= to; i++ {
		*result = append(*result, matrix[i][col])
	}
}

func runRL(matrix [][]int, row, from, to int, result *[]int) {
	for i := from; i >= to; i-- {
		*result = append(*result, matrix[row][i])
	}
}

func runDU(matrix [][]int, col, from, to int, result *[]int) {
	for i := from; i >= to; i-- {
		*result = append(*result, matrix[i][col])
	}
}
