package solution

import "math"

func TopDown(s string) int {
	if len(s) == 0 {
		return 0
	}

	cache := map[int]map[int]map[int]int{}
	return recurse(s, 0, 0, 0, 0, cache)
}

func recurse(
	s string,
	index int,
	openPar, closePar int,
	validLen int,
	cache map[int]map[int]map[int]int,
) int {
	if index == len(s) {
		if openPar == closePar {
			return validLen
		}

		return 0
	}

	if openPar < closePar {
		return 0
	}

	if _, ok := cache[index]; ok {
		if _, ok := cache[index][openPar]; ok {
			if result, ok := cache[index][openPar][closePar]; ok {
				return result
			}
		}
	}

	var max float64 = 0

	// Not skip
	if s[index] == '(' {
		max = float64(recurse(
			s,
			index+1,
			openPar+1,
			closePar,
			validLen+1,
			cache,
		))
	} else if s[index] == ')' {
		max = float64(recurse(
			s,
			index+1,
			openPar,
			closePar+1,
			validLen+1,
			cache,
		))
	}

	// Skip
	skip := recurse(
		s,
		index+1,
		openPar,
		closePar,
		validLen,
		cache,
	)
	max = math.Max(max, float64(skip))

	if _, ok := cache[index]; !ok {
		cache[index] = map[int]map[int]int{}
	}

	if _, ok := cache[index][openPar]; !ok {
		cache[index][openPar] = map[int]int{}
	}

	cache[index][openPar][closePar] = int(max)

	return int(max)
}
