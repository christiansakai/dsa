package solution

func Solve(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return Solve(n-1) + Solve(n-2)
}
