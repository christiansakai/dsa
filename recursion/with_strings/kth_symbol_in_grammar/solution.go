package solution

func Solve(N int, K int) int {
	if N == 1 && K == 1 {
		return 0
	}

	if K%2 == 0 {
		subProb := Solve(N-1, K/2)
		if subProb == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		subProb := Solve(N-1, (K+1)/2)
		if subProb == 0 {
			return 0
		} else {
			return 1
		}
	}
}
