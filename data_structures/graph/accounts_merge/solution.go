package solution

func Solve(accounts [][]string) [][]string {
	dict := map[string][]int{}

	for i, acc := range accounts {
		for j := 1; j < len(acc); j++ {
			mail := acc[j]

			if _, ok := dict[mail]; !ok {
				dict[mail] = []int{}
			}

			dict[mail] = append(dict[mail], i)
		}
	}

	merged := [][]string{}

	for mail, index := range dict {
		name := accounts[index][0]

	}

	return merged
}
