package solution

func Solve(n int) string {
	return string(recurse(n))
}

func recurse(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}

	if n == 1 {
		return []byte{'1'}
	}

	remainder := n % 2
	subProb := recurse(n / 2)

	if remainder == 0 {
		subProb = append(subProb, '0')
	} else {
		subProb = append(subProb, '1')
	}

	return subProb
}
