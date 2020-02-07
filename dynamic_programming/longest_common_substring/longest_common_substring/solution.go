package solution

import "math"

func TopDown(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}

	cache := map[int]map[int]map[int]int{}

	return recurse([]byte(s1), []byte(s2), 0, 0, 0, cache)
}

func recurse(s1, s2 []byte, i1, i2, length int, cache map[int]map[int]map[int]int) int {
	if i1 == len(s1) || i2 == len(s2) {
		return length
	}

	if _, ok := cache[i1]; ok {
		if _, ok := cache[i1][i2]; ok {
			if result, ok := cache[i1][i2][length]; ok {
				return result
			}
		}
	}

	if s1[i1] != s2[i2] {
		max := float64(length)

		skipS1 := recurse(s1, s2, i1+1, i2, 0, cache)
		max = math.Max(max, float64(skipS1))

		skipS2 := recurse(s1, s2, i1, i2+1, 0, cache)
		max = math.Max(max, float64(skipS2))

		result := int(max)

		if _, ok := cache[i1]; !ok {
			cache[i1] = map[int]map[int]int{}
		}

		if _, ok := cache[i1][i2]; !ok {
			cache[i1][i2] = map[int]int{}
		}

		cache[i1][i2][length] = result

		return result
	}

	max := float64(length)

	bothSame := recurse(s1, s2, i1+1, i2+1, length+1, cache)
	max = math.Max(max, float64(bothSame))

	skipS1 := recurse(s1, s2, i1+1, i2, 0, cache)
	max = math.Max(max, float64(skipS1))

	skipS2 := recurse(s1, s2, i1, i2+1, 0, cache)
	max = math.Max(max, float64(skipS2))

	result := int(max)

	if _, ok := cache[i1]; !ok {
		cache[i1] = map[int]map[int]int{}
	}

	if _, ok := cache[i1][i2]; !ok {
		cache[i1][i2] = map[int]int{}
	}

	cache[i1][i2][length] = result

	return result
}
