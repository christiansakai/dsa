package solution

func Solve(s string) string {
	if len(s) <= 1 {
		return s
	}

	b := []byte(s)

	i := 0
	for i < len(b)-1 {
		j := i
		for ; j < len(b)-1 && b[j+1] != ' '; j++ {
		}

		reverse(b, i, j)

		i = j + 1 + 1 // Skip space
	}

	return string(b)
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
