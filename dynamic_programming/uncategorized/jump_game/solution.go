package solution

func BottomUp(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	reachable := make([]bool, len(nums))
	reachable[0] = true

	for n := 1; n < len(nums); n++ {
		for i := 0; i < n; i++ {
			reachable[n] = reachable[n] || ((nums[i] >= (n - i)) && reachable[i])
		}
	}

	return reachable[len(nums)-1]
}
