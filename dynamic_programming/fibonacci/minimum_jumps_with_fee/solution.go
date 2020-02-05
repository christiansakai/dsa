package solution

import "math"

func TopDown(stairs []int) int {
	if len(stairs) == 0 {
		return 0
	}

	cache := map[int]int{}

	return recurse(stairs, len(stairs), cache)
}

func recurse(stairs []int, index int, cache map[int]int) int {
	if index == 0 {
		return 0
	}

	if index == 1 {
		return stairs[0]
	}

	if index == 2 {
		return int(math.Min(float64(stairs[0]), float64(stairs[0]+stairs[1])))
	}

	if result, ok := cache[index]; ok {
		return result
	}

	var min float64 = math.Inf(1)

	prevOne := stairs[index-1] + recurse(stairs, index-1, cache)
	min = math.Min(min, float64(prevOne))

	prevTwo := stairs[index-2] + recurse(stairs, index-2, cache)
	min = math.Min(min, float64(prevTwo))

	prevThree := stairs[index-3] + recurse(stairs, index-3, cache)
	min = math.Min(min, float64(prevThree))

	result := int(min)

	cache[index] = result

	return result
}
