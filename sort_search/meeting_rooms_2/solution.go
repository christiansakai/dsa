package solution

import (
	"fmt"
	"math"
	"sort"
)

func Solve(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sorted := ByStart(intervals)
	sort.Sort(sorted)

	meetingRooms := map[int][][]int{}
	roomNum := 0

	for i, interval := range intervals {
		if i == 0 {
			roomNum += 1
			meetingRooms[roomNum] = [][]int{interval}
		} else {
			intervalIsAddedToRoom := false

			for roomNum, intervals := range meetingRooms {
				lastIntervalInTheRoom := intervals[len(intervals)-1]

				if !isOverlap(interval, lastIntervalInTheRoom) {
					meetingRooms[roomNum] = append(meetingRooms[roomNum], interval)
					intervalIsAddedToRoom = true
					break
				}
			}

			if !intervalIsAddedToRoom {
				roomNum += 1
				meetingRooms[roomNum] = [][]int{interval}
			}
		}
	}

	fmt.Println(meetingRooms)

	return roomNum
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

func isOverlap(intervalA, intervalB []int) bool {
	leftBound := math.Max(float64(intervalA[0]), float64(intervalB[0]))
	rightBound := math.Min(float64(intervalB[1]), float64(intervalB[1]))

	return (rightBound - leftBound) > 0
}
