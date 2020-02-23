package solution

import "sort"

func Solve(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Sort(increasing(nums))

	result := [][]int{}

	i := 0
	for i < len(nums)-2 {
		j := i + 1
		for j < len(nums)-1 {
			remaining := 0 - (nums[i] + nums[j])

			k := binarySearch(nums, j+1, remaining)

			if k != -1 {
				result = append(result, []int{
					nums[i], nums[j], nums[k],
				})
			}

			for ; j+1 < len(nums)-1 && nums[j] == nums[j+1]; j++ {
			}
			j++
		}

		for ; i+1 < len(nums)-2 && nums[i] == nums[i+1]; i++ {
		}
		i++
	}

	return result
}

type increasing []int

func (a increasing) Len() int {
	return len(a)
}

func (a increasing) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a increasing) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func binarySearch(nums []int, start, target int) int {
	left := start
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
