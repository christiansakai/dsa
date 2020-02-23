package solution

func Solve(n int, guess func(int) int) int {
	left := 0
	right := n

	for left <= right {
		mid := left + (right-left)/2
		called := guess(mid)

		if called == 0 {
			return mid
		} else if called == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
