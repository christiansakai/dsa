package solution

func Solve(s, t string) bool {
	dict := map[rune]int{}

	for _, ch := range s {
		if _, ok := dict[ch]; ok {
			dict[ch] += 1
		} else {
			dict[ch] = 1
		}
	}

	for _, ch := range t {
		if count, ok := dict[ch]; ok {
			if count == 1 {
				delete(dict, ch)
			} else {
				dict[ch] -= 1
			}
		} else {
			return false
		}
	}

	return len(dict) == 0
}
