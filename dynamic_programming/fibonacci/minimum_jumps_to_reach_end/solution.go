package solution

import "math"

func TopDown(jumps []int) int {
	if len(jumps) == 0 {
		return 0
	}

	return recurse(jumps, len(jumps)-1)
}

func recurse(jumps []int, index int) int {
	if index == 0 {
		return 0
	}

	var min float64 = math.Inf(1)
	for i := 0; i < index; i++ {
		if jumps[i] >= index-i {
			from := 1 + recurse(jumps, i)
			min = math.Min(min, float64(from))
		}
	}

	return int(min)
}
