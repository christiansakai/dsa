package solution

// (n - 1)^2 = n^2 - 2n + 1
// n^2 = (n - 1)^2 + 2n - 1
func Solve(n int) int {
	if n == 0 {
		return 0
	}

	return Solve(n-1) + (2 * n) - 1
}
