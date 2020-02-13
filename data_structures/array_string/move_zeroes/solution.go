package solution

import "fmt"

func Solve(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && i > 0 {
			j := i

			for ; j > 0; j-- {
				if nums[j-1] != 0 {
					break
				}
			}

			fmt.Println(i, j)

			if i != j {
				nums[j] = nums[i]
				nums[i] = 0
			}
		}
	}
}
