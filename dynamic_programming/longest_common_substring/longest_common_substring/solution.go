package solution

// import "math"

func TopDown(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	return recurse([]byte(s1), []byte(s2), 0, 0)
}

func recurse(s1, s2 []byte, i1, i2 int) int {
	if i1 >= len(s1) || i2 >= len(s2) {
		return 0
	}

	if s1[i1] == s2[i2] {
		return 1 + recurse(s1, s2, i1+1, i2+1)
	}

	// return recurse(s1, s2, i1 + 1, i2 + 1)

	// nextI1 := recurse(s1, s2, i1 + 1, i2)
	// nextI2 := recurse(s1, s2, i1, i2 + 1)

	// max := math.Max(float64(nextI1), float64(nextI2))

	// return int(max)
}
