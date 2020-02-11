package solution

func Solve(str string) int {
	if len(str) == 0 {
		return 0
	}

	return recurse([]byte(str), 0)
}

func recurse(str []byte, index int) int {
	if len(str) == index {
		return 0
	}

	if isInt(str[index]) {
		integer := toInt(str[index])
		return integer + recurse(str, index+1)
	}

	return recurse(str, index+1)
}

func isInt(b byte) bool {
	return b == byte('0') ||
		b == byte('1') ||
		b == byte('2') ||
		b == byte('3') ||
		b == byte('4') ||
		b == byte('5') ||
		b == byte('6') ||
		b == byte('7') ||
		b == byte('8') ||
		b == byte('9')
}

func toInt(b byte) int {
	switch {
	case b == byte('0'):
		return 0
	case b == byte('1'):
		return 1
	case b == byte('2'):
		return 2
	case b == byte('3'):
		return 3
	case b == byte('4'):
		return 4
	case b == byte('5'):
		return 5
	case b == byte('6'):
		return 6
	case b == byte('7'):
		return 7
	case b == byte('8'):
		return 8
	case b == byte('9'):
		return 9
	}

	return -1
}
