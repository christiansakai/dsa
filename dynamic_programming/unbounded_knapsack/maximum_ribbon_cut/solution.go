package solution

import "math"

func TopDown(cuttings []int, length int) int {
	if len(cuttings) == 0 {
		return 0
	}

	return recurse(cuttings, length)
}

func recurse(cuttings []int, length int) int {
	if length == 0 {
		return 0
	}

	var max float64 = 0
	for _, c := range cuttings {
		if c <= length {
			selection := 1 + recurse(cuttings, length-c)
			max = math.Max(max, float64(selection))
		}
	}

	result := int(max)

	return result
}

func BottomUp(cuttings []int, length int) int {
	if len(cuttings) == 0 {
		return 0
	}

	maxCuttings := make([]int, length+1)
	maxCuttings[0] = 0

	for l := 1; l <= length; l++ {
		var max float64 = 0
		for _, c := range cuttings {
			if c <= l {
				selection := 1 + maxCuttings[l-c]
				max = math.Max(max, float64(selection))
			}
		}

		maxCuttings[l] = int(max)
	}

	return maxCuttings[length]
}

func ident(indent int) string {
	str := ""
	for i := 0; i < indent; i++ {
		str += "  "
	}

	return str
}
