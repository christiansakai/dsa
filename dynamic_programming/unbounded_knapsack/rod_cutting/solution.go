package solution

import "math"

type Cutting struct {
	Length int
	Price  int
}

func TopDown(cuttings []Cutting, length int) int {
	if length == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(cuttings, length, len(cuttings)-1, cache)
}

func recurse(cuttings []Cutting, length, index int, cache map[int]map[int]int) int {
	if length == 0 || index == -1 {
		return 0
	}

	if _, ok := cache[length]; ok {
		if result, ok := cache[length][index]; ok {
			return result
		}
	}

	var max float64 = 0

	if cuttings[index].Length <= length {
		with := cuttings[index].Price + recurse(cuttings, length-cuttings[index].Length, index, cache)
		max = math.Max(max, float64(with))
	}

	without := recurse(cuttings, length, index-1, cache)
	max = math.Max(max, float64(without))

	result := int(max)

	if _, ok := cache[length]; !ok {
		cache[length] = map[int]int{}
	}

	cache[length][index] = result

	return result
}
