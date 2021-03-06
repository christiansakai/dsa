package solution

func TopDown(n int) int {
	cache := map[int]int{}

	return recurse(n, cache)
}

func recurse(n int, cache map[int]int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if result, ok := cache[n]; ok {
		return result
	}

	result := recurse(n-1, cache) + recurse(n-2, cache)

	cache[n] = result

	return result
}
