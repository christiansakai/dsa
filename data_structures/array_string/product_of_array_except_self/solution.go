package solution

func Solve(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	productToTheLeft := make([]int, len(nums))
	productToTheLeft[0] = 1
	for i := 1; i < len(nums); i++ {
		productToTheLeft[i] = nums[i-1] * productToTheLeft[i-1]
	}

	productToTheRight := make([]int, len(nums))
	productToTheRight[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		productToTheRight[i] = nums[i+1] * productToTheRight[i+1]
	}

	productExceptSelf := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		productExceptSelf[i] = productToTheLeft[i] * productToTheRight[i]
	}

	return productExceptSelf
}
