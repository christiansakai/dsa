package solution

import "math"

func Solve(arr []int, s int) int {
	var min float64 = math.Inf(1)

	for i := 0; i < len(arr); i++ {
		sum := 0
		foundGreater := false

		j := i
		for ; j < len(arr); j++ {
			sum += arr[j]

			if sum >= s {
				foundGreater = true
				break
			}
		}

		if foundGreater {
			length := j - i + 1
			min = math.Min(min, float64(length))
		}
	}

	return int(min)
}
