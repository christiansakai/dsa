package solution

func Solve(nums []int) []int {
	qSort(nums, 0, len(nums)-1)
	return nums
}

func qSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		qSort(nums, left, p-1)
		qSort(nums, p+1, right)
	}
}

func partition(nums []int, left, right int) int {
	// picks last element as pivot
	// returns the index of pivot value in the sorted array
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			temp := nums[i]
			nums[i] = nums[j]
			i++

			nums[j] = temp
		}
	}

	temp := nums[i]
	nums[i] = nums[right]
	nums[right] = temp

	return i
}
