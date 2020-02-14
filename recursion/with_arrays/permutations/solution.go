package solution

func Solve(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	result := &[][]int{}
	set := &[]int{}

	recurse(nums, set, result)

	return *result
}

func recurse(nums []int, permutation *[]int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, createCopy(*permutation))
		return
	}

	for i := 0; i < len(nums); i++ {
		newNums := generateNewNums(nums, i)

		// Push the current element
		*permutation = append(*permutation, nums[i])

		recurse(newNums, permutation, result)

		// Pop the current element for cleanup
		*permutation = (*permutation)[:len(*permutation)-1]
	}
}

func generateNewNums(nums []int, excludeIndex int) []int {
	newNums := []int{}
	for i, el := range nums {
		if i != excludeIndex {
			newNums = append(newNums, el)
		}
	}

	return newNums
}

func createCopy(nums []int) []int {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[i] = nums[i]
	}

	return newNums
}
