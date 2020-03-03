package solution

func Solve(board [][]byte) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == '.' && recurse(board, r, c) {
				return
			}
		}
	}
}

func recurse(board [][]byte, row, col int) bool {
	if row == len(board) && row == len(board[0]) {
		return true
	}

	if row == len(board) {
		return recurse(board, 0, col+1)
	}

	if col == len(board[0]) {
		return recurse(board, row+1, 0)
	}

	for _, d := range digits {
		if !isValid(board, row, col, d) {
			return true
		}
	}

	return false
}

var digits = []byte{
	'1', '2', '3',
	'4', '5', '6',
	'7', '8', '9',
}

func isValid(board [][]byte, row, col int, digit byte) bool {
	return checkVertical(board, row, col, digit) &&
		checkHorizontal(board, row, col, digit) &&
		checkThreeByThree(board, row, col, digit)
}

func checkVertical(board [][]byte, row, col int, digit byte) bool {
	for r := 0; r < len(board); r++ {
		if row != r && board[r][col] == digit {
			return false
		}
	}

	return true
}

func checkHorizontal(board [][]byte, row, col int, digit byte) bool {
	for c := 0; c < len(board[0]); c++ {
		if col != c && board[row][c] == digit {
			return false
		}
	}

	return true
}

func checkThreeByThree(board, row, col, digit) bool {
	row := getStart(row)
	col := getStart(col)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := row + i
			c := col + j

			if !(row == c && col == c) && board[r][c] == digit {
				return false
			}
		}
	}

	return true
}

func getStart(num int) int {
	if num <= 2 {
		return 0
	}

	if num <= 5 {
		return 3
	}

	return 6
}
