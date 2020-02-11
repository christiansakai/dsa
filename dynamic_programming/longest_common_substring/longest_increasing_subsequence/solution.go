package solution

import "math"

// import "fmt"

func TopDown(sequence []int) int {
	if len(sequence) <= 1 {
		return len(sequence)
	}

	return recurse(sequence, len(sequence), len(sequence)-1)
}

func recurse(sequence []int, i, j int) int {
	// fmt.Println(i, j)
	if i == 0 || j == 0 {
		return 0
	}

	var max float64 = 0
	if sequence[i] > sequence[j] {
		with := 1 + recurse(sequence, i, j-1)
		max = math.Max(max, float64(with))
	}

	without := recurse(sequence, i-1, j)
	max = math.Max(max, float64(without))

	return int(max)
}
