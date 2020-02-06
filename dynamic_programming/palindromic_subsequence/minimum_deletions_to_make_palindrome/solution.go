package solution

import "math"

func SolveTopDown(str string) int {
	if len(str) == 0 {
		return 0
	}

  cache := map[int]map[int]int{}
	byteStr := []byte(str)
	lpsLength := recurse(byteStr, 0, len(byteStr)-1, cache)

	return len(str) - lpsLength
}

func recurse(str []byte, start, end int, cache map[int]map[int]int) int {
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

	var max float64 = 0

	if str[start] == str[end] {
		withFrontBack := 2 + recurse(str, start+1, end-1, cache)
		max = math.Max(max, float64(withFrontBack))
	}

	skipFront := recurse(str, start+1, end, cache)
	max = math.Max(max, float64(skipFront))

	skipBack := recurse(str, start, end-1, cache)
	max = math.Max(max, float64(skipBack))

	result := int(max)

  if _, ok := cache[start]; !ok {
    cache[start] = map[int]int{}
  }

  cache[start][end] = result

	return result
}
