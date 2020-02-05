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

	cache := map[int]int{}

	return recurse(cuttings, length, cache)
}

func recurse(cuttings []Cutting, length int, cache map[int]int) int {
	if length == 0 {
		return 0
	}

	if result, ok := cache[length]; ok {
		return result
	}

	var max float64 = 0
	for _, c := range cuttings {
		if c.Length <= length {
			selection := c.Price + recurse(cuttings, length-c.Length, cache)
			max = math.Max(max, float64(selection))
		}
	}

	result := int(max)

	cache[length] = result

	return result
}

func BottomUp(cuttings []Cutting, length int) int {
	if length == 0 {
		return 0
	}

	profits := make([]int, length+1)
	profits[0] = 0

	for l := 1; l <= length; l++ {
		var max float64 = 0
		for _, c := range cuttings {
			if c.Length <= l {
				selection := c.Price + profits[l-c.Length]
				max = math.Max(max, float64(selection))
			}
		}

		profits[l] = int(max)
	}

	return profits[length]
}
