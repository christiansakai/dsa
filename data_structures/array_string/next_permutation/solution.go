package solution

import "fmt"

func Solve(nums []int) {
	if len(nums) == 0 {
		return
	}

	decreasingSubsequenceStartIndex := findDecreasingSubsequenceStartIndex(nums)
	if decreasingSubsequenceStartIndex == 0 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	fmt.Println(decreasingSubsequenceStartIndex)

	needToBeSwapped := decreasingSubsequenceStartIndex - 1

	nextBiggerElementIndex := findNextBiggerElementFromRight(
		nums,
		decreasingSubsequenceStartIndex,
		nums[needToBeSwapped],
	)

	if nextBiggerElementIndex != -1 {
		swap(nums, needToBeSwapped, nextBiggerElementIndex)
		reverse(nums, needToBeSwapped+1, len(nums)-1)
	}
}

func findDecreasingSubsequenceStartIndex(nums []int) int {
	i := len(nums) - 1
	for ; i > 0 && (nums[i-1] >= nums[i]); i-- {
	}

	return i
}

func findNextBiggerElementFromRight(nums []int, start, element int) int {
	for i := len(nums) - 1; i >= start; i-- {
		if nums[i] > element {
			return i
		}
	}

	return -1
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start, end int) {
	for start < end {
		swap(nums, start, end)

		start++
		end--
	}
}
