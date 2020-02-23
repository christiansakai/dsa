package solution

import (
	"strconv"
	// "fmt"
	"strings"
)

func Solve(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	hashMap := map[string][]int{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			remainder := 0 - (nums[i] + nums[j])

			k := binarySearch(nums, j+1, len(nums)-1, remainder)
			if k != -1 {
				threeSum := []int{
					nums[i],
					nums[j],
					nums[k],
				}

				hashedStr := toString(threeSum)
				hashMap[hashedStr] = threeSum
			}
		}
	}

	result := [][]int{}
	for _, v := range hashMap {
		result = append(result, v)
	}

	return result
}

func binarySearch(nums []int, from, to, target int) int {
	left := from
	right := to

	for left <= right {
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

func toString(triplets []int) string {
	b := []string{}
	for _, el := range triplets {
		b = append(b, strconv.Itoa(el))
	}

	return strings.Join(b, "#")
}
