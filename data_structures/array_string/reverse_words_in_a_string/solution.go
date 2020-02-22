package solution

func Solve(s string) string {
	if len(s) <= 1 {
		return s
	}

	result := []byte{}
	b := []byte(s)

	i := 0
	for i < len(b)-1 {
		j := i
		for ; j < len(b)-1 && b[j+1] != ' '; j++ {
		}

		taken := take(b, i, j)
		taken = append(taken, ' ')

		result = append(taken, result...)

		return string(result)

		// j = j + 1
		// for ; j < len(b) - 1 && b[j + 1] == ' '; j++ {}

		// i = j + 1
	}

	return string(result)
}

func take(b []byte, from, to int) []byte {
	copied := make([]byte, to-from+1)
	for i := from; i <= to; i++ {
		copied[i] = b[i]
	}

	return copied
}

func reverse(b []byte, from, to int) {
	for from < to {
		temp := b[from]
		b[from] = b[to]
		b[to] = temp

		from += 1
		to -= 1
	}
}
