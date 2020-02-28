package solution

func Solve(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		celebCandidate := 0
		for i := 1; i < n; i++ {
			if knows(celebCandidate, i) {
				celebCandidate = i
			}
		}

		if isIndeedACeleb(n, knows, celebCandidate) {
			return celebCandidate
		}

		return -1
	}
}

func isIndeedACeleb(people int, knows func(a, b int) bool, candidate int) bool {
	for i := 0; i < people; i++ {
		if i != candidate {
			if knows(candidate, i) || !knows(i, candidate) {
				return false
			}
		}
	}

	return true
}
