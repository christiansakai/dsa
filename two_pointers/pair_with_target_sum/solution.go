package solution

func Solve(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	dict := map[int]int{}
	for i, el := range arr {
		remainder := target - el
		if remainderIndex, ok := dict[remainder]; ok {
			return []int{remainderIndex, i}
		}

		dict[el] = i
	}

	return []int{}
}
