package solution

func Solve(n int) [][]string {
	result := [][]string{}

	for row := 0; row < n; r++ {
		board := createEmptyBoard(n)
		recurse(n, row, 0, board, &result)
	}

	return result
}

func recurse(n, row, col int, board [][]byte, result *[][]string) {
	if col == n {
		result = append(result, toStringArr(board))
		return
	}

	if row == n {
		recurse(n, 0, col+1, board, result)
		return
	}

	if !isValid(row, col, board) {
		recurse(n, row+1, col, board, result)
	} else {
		board[row][col] = 'Q'
		recurse(n, row, col+1, board, result)
	}
}

func createEmptyBoard(n int) [][]byte {
	board := make([][]byte, n)
	for r := 0; r < n; r++ {
		board[r] = make([]byte, n)
	}

	return board
}

func toStringArr(board [][]byte) []string {
	result := make([]string, len(board))
	for r := 0; r < len(result); r++ {
		result[r] = string(board[r])
	}

	return result
}

func isValid(row, col int, board [][]byte) bool {
	// check horizontal
	for c := 0; c < len(board[0]); c++ {
		if c != col && board[row][c] == '.' {
			return false
		}
	}

	// check vertical
	for r := 0; r < len(board); r++ {
		if r != row && board[r][col] == '.' {
			return false
		}
	}

	// diagonal
	for r := 0; r < len(board); r++ {
	}
	for i := 0; i < len(board); i++ {
		otherRow := row + i
		if r != row && board[r][col] == '.' {
			return false
		}
	}

	return true
}
