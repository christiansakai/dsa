package solution

import "math"

func Solve(arr []int, s int) int {
	var min float64 = math.Inf(1)

	i := 0
	for i < len(arr) {
		j := i
		sum := 0

		for ; j < len(arr) && sum < s; j++ {
			sum += arr[j]
		}

		if sum >= s {
			min = math.Min(min, float64(j-i+1))

			for ; i < len(arr) && sum >= s; i++ {
				sum -= arr[i]
			}
		}
	}

	return int(min)
}
