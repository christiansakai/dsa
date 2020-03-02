package solution

func Solve(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		// If the mid element is on descending slope
		if nums[mid] > nums[mid+1] {
			// then the peak must be on the left
			right = mid
		} else { // If the mid element is on the ascending slope
			// then the peak must be on the right
			left = mid + 1
		}
	}

	return left
}
