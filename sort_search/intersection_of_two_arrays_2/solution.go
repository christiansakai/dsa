package solution

import "sort"

func Solve(nums1 []int, nums2 []int) []int {
	result := []int{}

	sort.Sort(increasing(nums1))
	sort.Sort(increasing(nums2))

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
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
