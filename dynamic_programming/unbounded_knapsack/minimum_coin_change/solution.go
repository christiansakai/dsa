package solution

import "math"

func TopDown(coins []int, total int) int {
	if len(coins) == 0 || total == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(coins, total, len(coins) - 1, cache)
}

func recurse(coins []int, total, index int, cache map[int]map[int]int) int {
	if total == 0 || index == -1 {
		return 0
	}

	// if _, ok := cache[total]; ok {
    // if result, ok := cache[total][index]; ok {
	// 	  return result
    // }
	// }

	var min float64 = math.Inf(1)

  if coins[index] <= total {
    with := 1 + recurse(coins, total - coins[index], index, cache)
    min = math.Min(min, float64(with))
  }

  without := recurse(coins, total, index - 1, cache)
  min = math.Min(min, float64(without))

	result := int(min)

  // if _, ok := cache[total]; !ok {
  //   cache[total] = map[int]int{}
  // }

	// cache[total][index] = result

	return result
}
