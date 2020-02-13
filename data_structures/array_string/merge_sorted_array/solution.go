package solution

func Solve(nums1 []int, m int, nums2 []int, n int) {
	total := m + n

	p1 := 0
	p2 := 0

	for i := 0; i < total; i++ {
		if p2 < n {
			if p1 > m-1 {
				nums1[p1] = nums2[p2]

				p1 += 1
				p2 += 1
			} else if nums1[p1] <= nums2[p2] {
				p1 += 1
			} else {
				shiftRight(nums1, p1, m)
				// after we shift we need to increase the cutoff for m
				m += 1

				nums1[p1] = nums2[p2]

				p1 += 1
				p2 += 1
			}
		}
	}
}

func shiftRight(nums []int, pos, limit int) {
	if pos > limit-1 {
		return
	}

	for i := len(nums) - 2; i >= pos; i-- {
		nums[i+1] = nums[i]
	}
}
