package solution

import "math"

func TopDown(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	cache := map[int]map[int]int{}

	return recurse([]byte(s1), []byte(s2), 0, 0, cache)
}

func recurse(s1, s2 []byte, i1, i2 int, cache map[int]map[int]int) int {
	if i1 >= len(s1) || i2 >= len(s2) {
		return 0
	}

	if _, ok := cache[i1]; ok {
		if result, ok := cache[i1][i2]; ok {
			return result
		}
	}

	var max float64 = 0
	if s1[i1] == s2[i2] {
		longestCommonSubStr := recurse(s1, s2, i1+1, i2+1, cache)
		longestCommonStr := 1 + longestCommonSubStr
		max = math.Max(max, float64(longestCommonStr))
	}

	skipS1 := recurse(s1, s2, i1+1, i2, cache)
	max = math.Max(max, float64(skipS1))

	skipS2 := recurse(s1, s2, i1, i2+1, cache)
	max = math.Max(max, float64(skipS2))

	result := int(max)

	if _, ok := cache[i1]; !ok {
		cache[i1] = map[int]int{}
	}

	cache[i1][i2] = result

	return result
}
