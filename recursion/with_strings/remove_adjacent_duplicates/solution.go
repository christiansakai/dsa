package solution

func Solve(str string) string {
	if len(str) <= 1 {
		return str
	}

	return string(recurse([]byte(str), 0))
}

func recurse(str []byte, index int) []byte {
	if index == len(str)-1 {
		return []byte{str[index]}
	}

	if str[index] == str[index+1] {
		return recurse(str, index+1)
	}

	newStr := []byte{str[index]}
	subProb := recurse(str, index+1)
	newStr = append(newStr, subProb...)

	return newStr
}
