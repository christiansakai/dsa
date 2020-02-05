package solution

import "math"

func TopDown(coins []int, total int) int {
	if len(coins) == 0 {
		return 0
	}

	cache := map[int]int{}

	return recurse(coins, total, cache)
}

func recurse(coins []int, total int, cache map[int]int) int {
	if total == 0 {
		return 0
	}

	if result, ok := cache[total]; ok {
		return result
	}

	var min float64 = math.Inf(1)
	for _, c := range coins {
		if c <= total {
			selection := 1 + recurse(coins, total-c, cache)
			min = math.Min(min, float64(selection))
		}
	}

	result := int(min)

	cache[total] = result

	return result
}

func BottomUp(coins []int, total int) int {
	if len(coins) == 0 {
		return 0
	}

	minCoins := make([]int, total+1)
	minCoins[0] = 0

	for t := 1; t <= total; t++ {
		var min float64 = math.Inf(1)
		for _, c := range coins {
			if c <= t {
				selection := 1 + minCoins[t-c]
				min = math.Min(min, float64(selection))
			}
		}

		minCoins[t] = int(min)
	}

	return minCoins[total]
}
