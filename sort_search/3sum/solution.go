package solution

import "sort"

func Solve(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Sort(increasing(nums))

	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		remainder := 0 - nums[i]

		points := getTwoSums(nums, i+1, remainder)

		for _, p := range points {
			result = append(result, []int{
				nums[i],
				nums[p.j],
				nums[p.k],
			})
		}

		for ; i < len(nums)-2 && nums[i] == nums[i+1]; i++ {
		}
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

type twoSum struct {
	j int
	k int
}

func getTwoSums(nums []int, start, target int) []twoSum {
	result := []twoSum{}

	left := start
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			result = append(result, twoSum{left, right})

			left += 1
			right -= 1

			for ; left < right && nums[left-1] == nums[left]; left++ {
			}
			for ; left < right && nums[right] == nums[right+1]; right-- {
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return result
}
