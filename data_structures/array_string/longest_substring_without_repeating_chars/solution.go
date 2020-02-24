package solution

import "math"

func Solve(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var max float64 = 0

	for i := 0; i < len(s); i++ {
		j := runUntilDuplicate(s, i)
		max = math.Max(max, float64(j-i))
	}

	return int(max)
}

func runUntilDuplicate(s string, index int) int {
	dict := map[byte]bool{}
	i := index
	for ; i < len(s); i++ {
		ch := s[i]

		if _, ok := dict[ch]; ok {
			return i
		}

		dict[ch] = true
	}

	return i
}
