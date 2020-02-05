package solution

import "math"

func SolveTopDown(str string) int {
	if len(str) == 0 {
		return 0
	}

	byteStr := []byte(str)
	lpsLength := recurse(byteStr, 0, len(byteStr)-1)

	return len(str) - lpsLength
}

func recurse(str []byte, start, end int) int {
	if start > end {
		return 0
	}

	if start == end {
		return 1
	}

	var max float64 = 0

	if str[start] == str[end] {
		withFrontBack := 2 + recurse(str, start+1, end-1)
		max = math.Max(max, float64(withFrontBack))
	}

	skipFront := recurse(str, start+1, end)
	max = math.Max(max, float64(skipFront))

	skipBack := recurse(str, start, end-1)
	max = math.Max(max, float64(skipBack))

	result := int(max)

	return result
}
