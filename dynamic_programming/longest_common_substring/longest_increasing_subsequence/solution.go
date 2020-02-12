package solution

import "math"

func TopDown(sequence []int) int {
	if len(sequence) <= 1 {
		return len(sequence)
	}

	cache := map[int]map[int]int{}

	return recurse(sequence, len(sequence), len(sequence)-1, cache)
}

func recurse(sequence []int, i, j int, cache map[int]map[int]int) int {
	if i == 0 || j == 0 {
		return 0
	}

	if _, ok := cache[i]; ok {
		if result, ok := cache[i][j]; ok {
			return result
		}
	}

	var max float64 = 0
	if sequence[i] > sequence[j] {
		with := 1 + recurse(sequence, i, j-1, cache)
		max = math.Max(max, float64(with))
	}

	without := recurse(sequence, i-1, j, cache)
	max = math.Max(max, float64(without))

	result := int(max)

	if _, ok := cache[i]; !ok {
		cache[i] = map[int]int{}
	}

	cache[i][j] = result

	return int(max)
}
