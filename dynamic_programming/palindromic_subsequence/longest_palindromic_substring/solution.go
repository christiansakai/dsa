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
	if start == end {
		return 1
	}

	var max float64 = 0

	if str[start] == str[end] {
		withFrontBack := 2 + recurse(str, start+1, end-1)
	}

}
