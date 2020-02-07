package solution

import "math"

func TopDown(str string) int {
	if len(str) == 0 {
		return 0
	}

	byteStr := []byte(str)

	return recurse(byteStr, 0, len(byteStr)-1)
}

func recurse(str []byte, start, end int) int {
	if start > end {
		return 0
	}

	var min float64 = math.Inf(1)

	if str[start] == str[end] {
		subProbMinCut := recurse(str, start+1, end-1)

		// If subProbMinCut == 0 that means subproblem is a palindrome
		if subProbMinCut == 0 {
			min = 0
		} else {
			// Else, it is not a palindrome so we cut left and right
			min = 2 + math.Min(min, float64(subProbMinCut))
		}
	}

	skipFront := 1 + recurse(str, start+1, end)
	min = math.Min(min, float64(skipFront))

	skipBack := 1 + recurse(str, start, end-1)
	min = math.Min(min, float64(skipBack))

	return int(min)
}
