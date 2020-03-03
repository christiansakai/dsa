package solution

func Solve(n int) []string {
	result := []string{}
	parentheses := []byte{}

	recurse(n, 0, 0, &parentheses, &result)

	return result
}

func recurse(total, openPar, closePar int, parentheses *[]byte, result *[]string) {
	if openPar < closePar {
		return
	}

	if openPar == total {
		if openPar > closePar {
			*parentheses = append(*parentheses, ')')
			recurse(total, openPar, closePar+1, parentheses, result)
			*parentheses = (*parentheses)[:len(*parentheses)-1]
		} else if openPar == closePar {
			*result = append(*result, toString(parentheses))
		}

		return
	}

	*parentheses = append(*parentheses, '(')
	recurse(total, openPar+1, closePar, parentheses, result)
	*parentheses = (*parentheses)[:len(*parentheses)-1]

	*parentheses = append(*parentheses, ')')
	recurse(total, openPar, closePar+1, parentheses, result)
	*parentheses = (*parentheses)[:len(*parentheses)-1]
}

func toString(parentheses *[]byte) string {
	result := make([]byte, len(*parentheses))
	for i := 0; i < len(*parentheses); i++ {
		result[i] = (*parentheses)[i]
	}

	return string(result)
}
