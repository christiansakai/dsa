package solution

func Solve(n int) []string {
	if n == 0 {
		return []string{}
	}

	digits := []byte{'0', '1', '6', '8', '9'}
	result := []string{}
	str := []byte{}

	recurse(digits, n, 0, *str, *result)

}
