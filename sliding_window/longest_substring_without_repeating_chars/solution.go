package solution

import "math"

func Solve(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var max float64 = 0
	dict := map[byte]bool{}

	i := 0
	j := 0

	for i < len(s) && j < len(s) {
		encounteredChar := s[j]

		if _, ok := dict[encounteredChar]; !ok {
			dict[encounteredChar] = true
			max = math.Max(max, float64(j-i+1))
			j += 1
		} else {
			forgottenChar := s[i]
			delete(dict, forgottenChar)
			i += 1
		}
	}

	return int(max)
}
