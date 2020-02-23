package solution

func Solve(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left+1 < right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	return -1
}
