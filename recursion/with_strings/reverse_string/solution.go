package solution

func Solve(s []byte) {
	if len(s) == 0 {
		return
	}

	recurse(s, 0)
}

func recurse(s []byte, index int) {
	if index == len(s)/2 {
		return
	}

	temp := s[index]
	s[index] = s[len(s)-1-index]
	s[len(s)-1-index] = temp

	recurse(s, index+1)
}
