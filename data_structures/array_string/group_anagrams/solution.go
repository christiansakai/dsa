package solution

func Solve(strs []string) [][]string {
	dict := map[string][]string{}

	for _, str := range strs {
		strHash := calculateStrHash(str)
		if _, ok := dict[strHash]; ok {
			dict[strHash] = append(dict[strHash], str)
		} else {
			dict[strHash] = []string{str}
		}
	}

	result := [][]string{}
	for _, v := range dict {
		result = append(result, v)
	}

	return result
}

func calculateStrHash(str string) string {
	charCount := make([]byte, 26)
	for _, ch := range str {
		index := ch - 97
		charCount[index] += 1
	}

	return string(charCount)
}
