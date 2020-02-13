package solution

func createRead4(file string) func([]byte) int {
	b := []byte(file)
	fp := 0

	return func(buf []byte) int {
		count := 0

		for i := 0; i < 4 && fp < len(b); i++ {
			buf[i] = b[fp]

			fp += 1
			count += 1
		}

		return count
	}
}
