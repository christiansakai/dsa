package solution

func Solve(s1, s2 string) string {
	if len(s1) == 0 {
		return s2
	}

	if len(s2) == 0 {
		return s1
	}

	return string(recurse([]byte(s1), []byte(s2), 0, 0))
}

func recurse(s1, s2 []byte, i1, i2 int) []byte {
	if i1 == len(s1) && i2 == len(s2) {
		return []byte{}
	}

	if i1 == len(s1) {
		return s2[i2:]
	}

	if i2 == len(s2) {
		return s1[i1:]
	}

	if s1[i1] < s2[i2] {
		b := []byte{s1[i1]}
		subProb := recurse(s1, s2, i1+1, i2)

		b = append(b, subProb...)
		return b
	}

	b := []byte{s2[i2]}
	subProb := recurse(s1, s2, i1, i2+1)

	b = append(b, subProb...)
	return b
}
