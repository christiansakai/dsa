package solution

import (
	"math"
	"sort"
)

func Solve(intervals [][]int) [][]int {
	sorted := ByStart(intervals)
	sort.Sort(sorted)

	result := [][]int{}

	for i, interval := range sorted {
		if i == 0 {
			result = append(result, interval)
		} else {
			lastInterval := result[len(result)-1]

			if isOverlap(lastInterval, interval) {
				merged := runMerge(lastInterval, interval)

				result = result[:len(result)-1]
				result = append(result, merged)
			} else {
				result = append(result, interval)
			}
		}
	}

	return result
}

type ByStart [][]int

func (b ByStart) Len() int {
	return len(b)
}

func (b ByStart) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

func (b ByStart) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func isOverlap(i1, i2 []int) bool {
	leftBound := int(math.Max(float64(i1[0]), float64(i2[0])))
	rightBound := int(math.Min(float64(i1[1]), float64(i2[1])))

	return (rightBound - leftBound) >= 0
}

func runMerge(i1, i2 []int) []int {
	leftBound := int(math.Min(float64(i1[0]), float64(i2[0])))
	rightBound := int(math.Max(float64(i1[1]), float64(i2[1])))

	return []int{leftBound, rightBound}
}
