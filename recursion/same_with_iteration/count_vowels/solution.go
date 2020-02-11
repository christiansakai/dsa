package solution

func Solve(str string) int {
	if len(str) == 0 {
		return 0
	}

	return recurse([]byte(str), 0)
}

func recurse(str []byte, index int) int {
	if index == len(str) {
		return 0
	}

	if isVowel(str[index]) {
		return 1 + recurse(str, index+1)
	}

	return recurse(str, index+1)
}

func isVowel(b byte) bool {
	return b == byte('a') ||
		b == byte('i') ||
		b == byte('u') ||
		b == byte('e') ||
		b == byte('o')
}
