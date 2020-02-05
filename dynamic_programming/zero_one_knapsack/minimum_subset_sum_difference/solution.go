package solution

import "math"

func TopDown(set []int) int {
	if len(set) == 0 {
		return 0
	}

	// We don't need to cache totalB because
	// totalB is basically a function of totalA
	cache := map[int]map[int]int{}

	result := recurse(set, len(set)-1, 0, 0, cache)

	return int(result)
}

func recurse(set []int, index, totalA, totalB int, cache map[int]map[int]int) int {
	if index == -1 {
		return int(math.Abs(float64(totalA - totalB)))
	}

	if _, ok := cache[index]; ok {
		if result, ok := cache[index][totalA]; ok {
			return result
		}
	}

	putToA := recurse(set, index-1, totalA+set[index], totalB, cache)
	putToB := recurse(set, index-1, totalA, totalB+set[index], cache)

	result := int(math.Min(float64(putToA), float64(putToB)))

	if _, ok := cache[index]; !ok {
		cache[index] = map[int]int{}
	}

	cache[index][totalA] = result

	return result
}
