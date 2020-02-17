package solution

func Solve(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}

	resultSet := map[string]bool{}
	str := []byte{}

	invalidOpen, invalidClose := findInvalidCount(s)
	validLenShoudBe := len(s) - invalidOpen - invalidClose

	recurse(s, 0, 0, validLenShoudBe, &str, resultSet)

	result := []string{}
	for k, _ := range resultSet {
		result = append(result, k)
	}

	return result
}

func recurse(
	s string,
	index, invalidOpen int,
	validLenShoudBe int,
	str *[]byte,
	result map[string]bool,
) {
	if index == len(s) {
		if invalidOpen == 0 && len(*str) == validLenShoudBe {
			copied := string(copyArr(*str))
			result[copied] = true
		}

		return
	}

	if s[index] == '(' {
		// Skip
		recurse(s, index+1, invalidOpen, validLenShoudBe, str, result)

		// Not skip
		*str = append(*str, s[index])
		recurse(s, index+1, invalidOpen+1, validLenShoudBe, str, result)
		*str = (*str)[:len(*str)-1]
	} else if s[index] == ')' {
		// Skip
		recurse(s, index+1, invalidOpen, validLenShoudBe, str, result)

		// Not skip
		*str = append(*str, s[index])
		if invalidOpen > 0 {
			recurse(s, index+1, invalidOpen-1, validLenShoudBe, str, result)
		}
		*str = (*str)[:len(*str)-1]
	} else {
		*str = append(*str, s[index])
		recurse(s, index+1, invalidOpen, validLenShoudBe, str, result)
		*str = (*str)[:len(*str)-1]
	}
}

func copyArr(arr []byte) []byte {
	newArr := make([]byte, len(arr))
	for i, el := range arr {
		newArr[i] = el
	}

	return newArr
}

func findInvalidCount(s string) (int, int) {
	invalidOpen := 0
	invalidClose := 0

	for _, ch := range s {
		if ch == '(' {
			invalidOpen += 1
		} else if ch == ')' {
			if invalidOpen > 0 {
				invalidOpen -= 1
			} else {
				invalidClose += 1
			}
		}
	}

	return invalidOpen, invalidClose
}
