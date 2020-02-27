package solution

func Solve(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == word[0] {
				if dfs(board, word, r, c) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, startRow, startCol int) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	return recurse(
		board,
		[]byte(word),
		0,
		startRow,
		startCol,
		visited,
	)
}

func recurse(
	board [][]byte,
	word []byte,
	index, row, col int,
	visited [][]bool,
) bool {
	visited[row][col] = true

	if index == len(word)-1 {
		if word[index] == board[row][col] {
			return true
		}

		return false
	}

	if word[index] != board[row][col] {
		return false
	}

	found := false

	// go up
	if row-1 >= 0 && !visited[row-1][col] {
		found = found || recurse(board, word, index+1, row-1, col, visited)
	}

	// go right
	if col+1 < len(board[0]) && !visited[row][col+1] {
		found = found || recurse(board, word, index+1, row, col+1, visited)
	}

	// go down
	if row+1 < len(board) && !visited[row+1][col] {
		found = found || recurse(board, word, index+1, row+1, col, visited)
	}

	// go left
	if col-1 >= 0 && !visited[row][col-1] {
		found = found || recurse(board, word, index+1, row, col-1, visited)
	}

	return found
}
