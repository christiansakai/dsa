package solution

func Solve(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	result := &[][]int{}
	subset := &[]int{}

	recurse(nums, 0, subset, result)

	return *result
}

func recurse(nums []int, index int, subset *[]int, result *[][]int) {
	if index == len(nums) {
		*result = append(*result, createCopy(*subset))
		return
	}

	// Without
	recurse(nums, index+1, subset, result)

	// With
	*subset = append(*subset, nums[index])
	recurse(nums, index+1, subset, result)
	*subset = (*subset)[:len(*subset)-1]
}

func createCopy(nums []int) []int {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[i] = nums[i]
	}

	return newNums
}
