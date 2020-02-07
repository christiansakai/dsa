package solution

import "math"

func TopDown(coins []int, total int) int {
	if len(coins) == 0 || total == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(coins, total, len(coins)-1, cache)
}

func recurse(coins []int, total, index int, cache map[int]map[int]int) int {
	if index == -1 {
		if total == 0 {
			return 0
		}

		// Signify that the total is not exhausted
		return int(math.Inf(1))
	}

	// Bubble to the top that total is not exhausted
	var min float64 = math.Inf(1)
	if coins[index] <= total {
		thisCoin := recurse(coins, total-coins[index], index, cache)

		if thisCoin != int(math.Inf(1)) {
			thisCoin += 1
			min = math.Min(min, float64(thisCoin))
		}
	}

	nextCoin := recurse(coins, total, index-1, cache)
	if nextCoin != int(math.Inf(1)) {
		min = math.Min(min, float64(nextCoin))
	}

	result := int(min)

	return result
}
