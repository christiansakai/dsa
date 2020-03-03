package solution

func Solve(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left, right := divide(nums)

	sortedLeft := Solve(left)
	sortedRight := Solve(right)

	return merge(sortedLeft, sortedRight)
}

func divide(nums []int) ([]int, []int) {
	pivot := (len(nums) / 2) - 1
	left := nums[:pivot+1]
	right := nums[pivot+1:]

	return left, right
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0

	l := 0
	r := 0

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] < right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}

		i++
	}

	return result
}
