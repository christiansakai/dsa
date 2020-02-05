package solution

import "math"

func TopDown(houses []int) int {
	if len(houses) == 0 {
		return 0
	}

	cache := map[int]int{}

	return recurse(houses, len(houses)-1, cache)
}

func recurse(houses []int, index int, cache map[int]int) int {
	if index == 0 {
		return houses[0]
	}

	if index == 1 {
		return houses[1]
	}

	with := houses[index] + recurse(houses, index-2, cache)
	without := recurse(houses, index-1, cache)

	result := int(math.Max(float64(with), float64(without)))

	return result
}
