package solution

import "math"

func SolveTopDown(str string) int {
	if len(str) == 0 {
		return 0
	}

	return int(recurse([]byte(str), 0, len(str)-1))
}

func recurse(str []byte, start, end int) float64 {
	if start > end {
		return 0
	}

	var total float64 = 0
	if str[start] == str[end] {
		total += 2
	}

	var max float64

	skipFront := recurse(str, start+1, end)
	skipBack := recurse(str, start, end-1)

	max = math.Max(max, skipFront)
	max = math.Max(max, skipBack)

	total += max

	return total
}
