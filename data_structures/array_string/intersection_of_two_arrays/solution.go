package solution

func Solve(nums1 []int, nums2 []int) []int {
	set := map[int]bool{}
	dict := map[int]int{}

	for _, el := range nums1 {
		if _, ok := dict[el]; ok {
			dict[el] += 1
		} else {
			dict[el] = 1
		}
	}

	for _, el := range nums2 {
		if count, ok := dict[el]; ok {
			if count > 1 {
				dict[el] -= 1
			} else {
				delete(dict, el)
			}

			set[el] = true
		}
	}

	result := []int{}
	for k, _ := range set {
		result = append(result, k)
	}

	return result
}
