package solution

func TopDown(stairs int) int {
	cache := map[int]int{}
	return recurse(stairs, cache)
}

func recurse(stairs int, cache map[int]int) int {
	if stairs == 1 {
		return 1
	}

	if stairs == 2 {
		return 2
	}

	if stairs == 3 {
		return 4
	}

	if result, ok := cache[stairs]; ok {
		return result
	}

	result := recurse(stairs-1, cache) + recurse(stairs-2, cache) + recurse(stairs-3, cache)

	cache[stairs] = result

	return result
}
