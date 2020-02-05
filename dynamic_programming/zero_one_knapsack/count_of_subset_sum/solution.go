package solution

func TopDown(set []int, S int) int {
	if len(set) == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(set, len(set)-1, S, cache)
}

func recurse(set []int, index, S int, cache map[int]map[int]int) int {
	if index == -1 {
		if S == 0 {
			return 1
		}

		return 0
	}

	if _, ok := cache[index]; ok {
		if result, ok := cache[index][S]; ok {
			return result
		}
	}

	with := recurse(set, index-1, S-set[index], cache)
	without := recurse(set, index-1, S, cache)

	result := with + without

	if _, ok := cache[index]; !ok {
		cache[index] = map[int]int{}
	}

	cache[index][S] = result

	return with + without
}
