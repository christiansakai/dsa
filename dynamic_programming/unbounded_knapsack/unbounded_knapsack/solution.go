package solution

import "math"

type Fruit struct {
	Name   string
	Weight int
	Profit int
}

func TopDown(fruits []Fruit, capacity int) int {
	if len(fruits) == 0 || capacity == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse(fruits, capacity, len(fruits)-1, cache)
}

func recurse(fruits []Fruit, capacity, index int, cache map[int]map[int]int) int {
	if capacity == 0 || index == -1 {
		return 0
	}

	if _, ok := cache[capacity]; ok {
		if result, ok := cache[capacity][index]; ok {
			return result
		}
	}

	var max float64

	if fruits[index].Weight <= capacity {
		with := fruits[index].Profit + recurse(fruits, capacity-fruits[index].Weight, index, cache)
		max = math.Max(max, float64(with))
	}

	without := recurse(fruits, capacity, index-1, cache)
	max = math.Max(max, float64(without))

	result := int(max)

	if _, ok := cache[capacity]; !ok {
		cache[capacity] = map[int]int{}
	}

	cache[capacity][index] = result

	return result
}
