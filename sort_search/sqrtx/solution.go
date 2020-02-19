package solution

import "fmt"

func Solve(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := left + (right-left)/2
		midSqr := mid * mid

		if midSqr == x {
		} else if midSqr < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
