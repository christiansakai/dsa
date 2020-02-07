package solution

import "math"

func TopDown(cuttings []int, length int) int {
	if len(cuttings) == 0 || length == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(cuttings, length, len(cuttings)-1, cache)
}

func recurse(cuttings []int, length, index int, cache map[int]map[int]int) int {
	if index == -1 {
		if length == 0 {
			return 0
		}

		// Signify that the length is not exhausted
		return int(math.Inf(-1))
	}

	if _, ok := cache[length]; ok {
		if result, ok := cache[length][index]; ok {
			return result
		}
	}

	// Bubble to the top that the length is not exhausted
	var max float64 = math.Inf(-1)

	if cuttings[index] <= length {
		thisCutting := recurse(cuttings, length-cuttings[index], index, cache)

		if thisCutting != int(math.Inf(-1)) {
			thisCutting += 1
			max = math.Max(max, float64(thisCutting))
		}
	}

	nextCutting := recurse(cuttings, length, index-1, cache)
	if nextCutting != int(math.Inf(-1)) {
		max = math.Max(max, float64(nextCutting))
	}

	result := int(max)

	if _, ok := cache[length]; !ok {
		cache[length] = map[int]int{}
	}

	cache[length][index] = result

	return result
}
