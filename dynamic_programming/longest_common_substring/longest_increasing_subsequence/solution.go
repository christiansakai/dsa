package solution

import "math"

func TopDown(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	cache := map[int]int{}

	return recurse(arr, 0, cache)
}

func recurse(arr []int, index int, cache map[int]int) int {
	if index == len(arr)-1 {
		return 1
	}

	if result, ok := cache[index]; ok {
		return result
	}

	var max float64 = 0
	if arr[index] < arr[index+1] {
		subProb := 1 + recurse(arr, index+1, cache)
		max = math.Max(max, float64(subProb))
	}

	nextI := recurse(arr, index+1, cache)
	max = math.Max(max, float64(nextI))

	result := int(max)

	cache[index] = result

	return result
}
