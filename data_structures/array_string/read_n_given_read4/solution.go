package solution

func Solve(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		temp := make([]byte, 4)
		bufPointer := 0
		nRemaining := n

		for nRemaining > 0 {
			count := read4(temp)
			if count == 0 {
				return bufPointer
			}

			for tempPointer := 0; tempPointer < count; tempPointer++ {
				buf[bufPointer] = temp[tempPointer]
				bufPointer += 1

				nRemaining -= 1
				if nRemaining == 0 {
					return bufPointer
				}
			}
		}

		return bufPointer
	}
}
