package solution

func Solve(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// If the first half of the array is NOT rotated
		if nums[left] <= nums[mid] {
			// Check if target is within the boundaries
			if nums[left] <= target && target < nums[mid] { // If so then indeed to left
				right = mid - 1
			} else { // Nope, go right
				left = mid + 1
			}
		} else { // If the second half of the array is NOT rotated
			// Check if target is within the boundaris
			if nums[mid] < target && target <= nums[right] { // If so then indeed go right
				left = mid + 1
			} else { // Nope, go left
				right = mid - 1
			}
		}
	}

	return -1
}
