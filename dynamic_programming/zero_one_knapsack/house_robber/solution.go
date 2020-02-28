package solution

import "math"

func TopDown(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cache := map[int]int{}
	return recurse(nums, len(nums)-1, cache)
}

func recurse(nums []int, index int, cache map[int]int) int {
	if index == -1 {
		return 0
	}

	if index == 0 {
		return nums[0]
	}

	if result, ok := cache[index]; ok {
		return result
	}

	robThisHouse := nums[index] + recurse(nums, index-2, cache)
	robNextHouse := recurse(nums, index-1, cache)

	result := int(math.Max(float64(robThisHouse), float64(robNextHouse)))

	cache[index] = result

	return result
}
