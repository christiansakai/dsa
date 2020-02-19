package solution

func TopDown(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		if s == "0" {
			return 0
		}

		return 1
	}

	return recurse(s, 0)
}

func recurse(s string, index int) int {
	if index == len(s) {
		return 1
	}

	total := 0
	if canGoFurther(s, index) {
		subProb := recurse(s, index+2)
		total += subProb
	}

	subProb := recurse(s, index+1)
	total += subProb

	return total
}

func canGoFurther(s string, index int) bool {
	return (s[index] == '1' || s[index] == '2') &&
		index+1 < len(s) &&
		(s[index+1] == '1' ||
			s[index+1] == '2' ||
			s[index+1] == '3' ||
			s[index+1] == '4' ||
			s[index+1] == '5' ||
			s[index+1] == '6')
}
