package solution

import "math"

func TopDown(s1, s2 string) int {
	if len(s1) == 0 && len(s2) == 0 {
		return 0
	}

	if len(s1) == 0 {
		return len(s2)
	}

	if len(s2) == 0 {
		return len(s1)
	}

	return recurse([]byte(s1), []byte(s2), 0, 0)
}

func recurse(s1, s2 []byte, i1, i2 int) int {
	if i1 == len(s1) && i2 == len(s2) {
		return 0
	}

	if i1 == len(s1) || i2 == len(s2) {
		return 1
	}

	var min float64 = math.Inf(1)

	if s1[i1] == s2[i2] {
		next := recurse(s1, s2, i1+1, i2+1)
		min = math.Min(min, float64(next))
	} else {
		next := 1 + recurse(s1, s2, i1+1, i2+1)
		min = math.Min(min, float64(next))
	}

	nextI1 := recurse(s1, s2, i1, i2+1)
	min = math.Min(min, float64(nextI1))

	nextI2 := recurse(s1, s2, i1+1, i2)
	min = math.Min(min, float64(nextI2))

	return int(min)
}
