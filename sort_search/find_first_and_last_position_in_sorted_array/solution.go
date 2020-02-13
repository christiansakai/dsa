package solution

func Solve(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	pos := binarySearch(nums, target)
	if pos == -1 {
		return []int{-1, -1}
	}

	left := pos
	right := pos

	for ; left >= 1; left-- {
		if nums[left-1] != target {
			break
		}
	}

	for ; right < len(nums)-1; right++ {
		if nums[right+1] != target {
			break
		}
	}

	return []int{left, right}
}

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for {
		if start > end {
			break
		}

		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
