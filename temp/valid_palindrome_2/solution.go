package solution

import "math"

func TopDown(s string) bool {
	if len(s) == 0 {
		return true
	}

	cache := map[int]map[int]int{}
	longestPalindromicSubsequence := lpsTopDown(s, 0, len(s)-1, cache)

	return len(s)-longestPalindromicSubsequence <= 1
}

func lpsTopDown(s string, start, end int, cache map[int]map[int]int) int {
	if start > end {
		return 0
	}

	if start == end {
		return 1
	}

	if _, ok := cache[start]; ok {
		if result, ok := cache[start][end]; ok {
			return result
		}
	}

	if s[start] == s[end] {
		return 2 + lpsTopDown(s, start+1, end-1, cache)
	}

	skipStart := lpsTopDown(s, start+1, end, cache)
	skipEnd := lpsTopDown(s, start, end-1, cache)

	max := int(math.Max(float64(skipStart), float64(skipEnd)))

	if _, ok := cache[start]; !ok {
		cache[start] = map[int]int{}
	}

	cache[start][end] = max

	return max
}
