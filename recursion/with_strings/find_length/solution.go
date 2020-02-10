package solution

func Solve(str string) int {
	if str == "" {
		return 0
	}

	return recurse([]byte(str), 0)
}

func recurse(str []byte, index int) int {
	if str[index] == 0 {
		return 0
	}

	return 1 + recurse(str, index+1)
}
