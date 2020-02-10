package solution

func Solve(row int) []int {
	if row == 0 {
		return []int{1}
	}

	prevRow := Solve(row - 1)

	nextRow := make([]int, len(prevRow)+1)
	nextRow[0] = 1

	for i := 1; i < len(prevRow); i++ {
		nextRow[i] = prevRow[i] + prevRow[i-1]
	}

	nextRow[len(nextRow)-1] = 1

	return nextRow
}
