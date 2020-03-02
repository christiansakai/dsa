package solution

func TopDown(str string) string {
	if len(str) == 0 {
		return ""
	}

	cache := map[int]map[int]strRange{}
	byteStr := []byte(str)

	palindromeRange := recurse(byteStr, 0, len(byteStr)-1, cache)

	return getString(byteStr, palindromeRange.from, palindromeRange.to)
}

type strRange struct {
	from int
	to   int
}

func recurse(str []byte, start, end int, cache map[int]map[int]strRange) strRange {
	if start > end {
		return strRange{start, end}
	}

	if start == end {
		return strRange{start, end}
	}

	if _, ok := cache[start]; ok {
		if result, ok := cache[start][end]; ok {
			return result
		}
	}

	if str[start] == str[end] {
		expectedPalindromicSubstringLength := end - start - 1

		palindromicSubstringRange := recurse(str, start+1, end-1, cache)
		palindromicSubstringLength := palindromicSubstringRange.to - palindromicSubstringRange.from + 1

		if expectedPalindromicSubstringLength == palindromicSubstringLength {
			return strRange{start, end}
		}
	}

	skipFrontRange := recurse(str, start+1, end, cache)
	skipFrontRangeLength := skipFrontRange.to - skipFrontRange.from + 1

	skipBackRange := recurse(str, start, end-1, cache)
	skipBackRangeLength := skipBackRange.to - skipBackRange.from + 1

	var result strRange

	if skipFrontRangeLength > skipBackRangeLength {
		result = skipFrontRange
	} else {
		result = skipBackRange
	}

	if _, ok := cache[start]; !ok {
		cache[start] = map[int]strRange{}
	}

	cache[start][end] = result

	return result
}

func getString(byteStr []byte, from, to int) string {
	result := []byte{}
	for i := from; i <= to; i++ {
		result = append(result, byteStr[i])
	}

	return string(result)
}
