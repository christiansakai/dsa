package solution

func TopDown(coins []int, total int) int {
	if len(coins) == 0 || total == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(coins, total, len(coins)-1, cache)
}

func recurse(coins []int, total, index int, cache map[int]map[int]int) int {
	if index == -1 {
		return 0
	}

	if total == 0 {
		return 1
	}

	if _, ok := cache[total]; ok {
		if result, ok := cache[total][index]; ok {
			return result
		}
	}

	ways := 0
	if coins[index] <= total {
		thisCoin := recurse(coins, total-coins[index], index, cache)
		ways += thisCoin
	}

	nextCoin := recurse(coins, total, index-1, cache)
	ways += nextCoin

	if _, ok := cache[total]; !ok {
		cache[total] = map[int]int{}
	}

	cache[total][index] = ways

	return ways
}
