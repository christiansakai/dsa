package solution

func Solve(str string) bool {
	if len(str) == 0 {
		return true
	}

	return recurse([]byte(str), 0, len(str)-1)
}

func recurse(str []byte, start, end int) bool {
	if start == end {
		return true
	}

	if str[start] != str[end] {
		return false
	}

	return true && recurse(str, start+1, end-1)
}
