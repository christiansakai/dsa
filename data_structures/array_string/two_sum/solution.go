package solution

func Solve(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	dict := map[int]int{}
	for i, n := range nums {
		remainder := target - n
		if j, ok := dict[remainder]; ok {
			return []int{j, i}
		}

		dict[n] = i
	}

	return []int{-1, -1}
}
