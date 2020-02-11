package solution

func Solve(str string) string {
	if len(str) == 0 {
		return ""
	}

	return string(recurse([]byte(str), 0))
}

func recurse(str []byte, index int) []byte {
	if index == len(str) {
		return []byte{}
	}

	if str[index] == '\t' || str[index] == ' ' {
		return recurse(str, index+1)
	}

	newStr := []byte{str[index]}
	subProb := recurse(str, index+1)
	newStr = append(newStr, subProb...)

	return newStr
}
