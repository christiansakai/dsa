package solution

import "fmt"

func Solve(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	if len(matrix) == 1 {
		return matrix[0]
	}

	result := []int{}

	elCount := len(matrix) * len(matrix[0])
	i := 0

	row := 0
	col := 0

	direction := "ur"

	for i < elCount {
		if direction == "ur" {
			takeCount, endRow, endCol := runUR(matrix, row, col, &result)
			i += takeCount

			if endRow == 0 && endCol == len(matrix[0])-1 {
				col = endCol
				row = endRow + 1
			} else if endRow == 0 {
				col = endCol + 1
				row = endRow
			} else {
				col = endCol
				row = endRow + 1
			}

			direction = "dl"
		} else {
			takeCount, endRow, endCol := runDL(matrix, row, col, &result)
			i += takeCount

			if endCol == 0 && endRow == len(matrix)-1 {
				col = endCol + 1
				row = endRow
			} else if endCol == 0 {
				col = endCol
				row = endRow + 1
			} else {
				col = endCol + 1
				row = endRow
			}

			direction = "ur"
		}

		fmt.Println(result)
	}

	return result
}

func runUR(matrix [][]int, row, col int, result *[]int) (int, int, int) {
	takeCount := 0

	for i := 0; ((row - i) >= 0) && (col+i < len(matrix[0])); i++ {
		r := row - i
		c := col + i

		takeCount += 1

		*result = append(*result, matrix[r][c])
	}

	return takeCount, row, col
}

func runDL(matrix [][]int, row, col int, result *[]int) (int, int, int) {
	takeCount := 0

	for i := 0; ((col - i) >= 0) && (row+i < len(matrix)); i++ {
		r := row + i
		c := col - i

		takeCount += 1

		*result = append(*result, matrix[r][c])
	}

	return takeCount, row, col
}
