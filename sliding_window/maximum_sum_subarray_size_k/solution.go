package solution

import "math"

// import "fmt"

func Solve(arr []int, k int) int {
	if len(arr) <= k {
		return sum(arr, 0, k-1)
	}

	var prevMax float64 = 0
	var max float64 = 0

	for i := 0; i <= len(arr)-k; i++ {
		if i == 0 {
			currMax := float64(sum(arr, i, i+k-1))
			max = currMax
			prevMax = currMax
		} else {
			currMax := float64(int(prevMax) - arr[i-1] + arr[i+k-1])
			prevMax = currMax
			max = math.Max(max, currMax)
		}
	}

	return int(max)
}

func sum(arr []int, from, to int) int {
	result := 0
	for i := from; i <= to; i++ {
		result += arr[i]
	}

	return result
}
