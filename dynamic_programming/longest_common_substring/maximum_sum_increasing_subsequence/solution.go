package solution

import "math"

func TopDown(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	return recurse(arr, 0)
}

func recurse(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[len(arr)-1]
	}

	var max float64 = 0
	if arr[index] < arr[index+1] {
		subProb := arr[index] + recurse(arr, index+1)
		max = math.Max(max, float64(subProb))
	}

	nextI := recurse(arr, index+1)
	max = math.Max(max, float64(nextI))

	return int(max)
}
