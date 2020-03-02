package solution

import "math"

func Solve(s string, k int) int {
	dict := map[byte]int{}
	var max float64 = 0

	i := 0
	for j := 0; j < len(s); j++ {
		ch := s[j]

		if _, ok := dict[ch]; ok {
			dict[ch] += 1
		} else {
			dict[ch] = 1
		}

		for countDistinctChars(dict) > k {
			ch := s[i]

			if dict[ch] > 1 {
				dict[ch] -= 1
			} else {
				delete(dict, ch)
			}

			i += 1
		}

		max = math.Max(max, float64(j-i+1))
	}

	return int(max)
}

func countDistinctChars(dict map[byte]int) int {
	count := 0
	for _, _ = range dict {
		count += 1
	}

	return count
}
