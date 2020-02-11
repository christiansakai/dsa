package solution

func Solve(n1, n2 int) int {
	if n1 == n2 {
		return n1
	}

	if n1 > n2 {
		return Solve(n1-n2, n2)
	}

	return Solve(n1, n2-n1)
}
