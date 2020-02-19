package solution

func Solve(nums []int) bool {
	dict := map[int]bool{}
	for _, el := range nums {
		if _, ok := dict[el]; ok {
			return true
		}

		dict[el] = true
	}

	return false
}
