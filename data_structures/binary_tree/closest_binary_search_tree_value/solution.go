package solution

import "math"

func Solve(root *node, target float64) int {
	return recurse(root, target, math.Inf(1), float64(-1))
}

func recurse(root *node, target, minDiff, closest float64) int {
	if root == nil {
		return int(closest)
	}

	rootValFl := float64(root.val)

	if rootValFl == target {
		return root.val
	} else if rootValFl < target {
		currDiff := math.Abs(rootValFl - target)
		if currDiff < minDiff {
			return recurse(root.right, target, currDiff, rootValFl)
		} else {
			return recurse(root.right, target, minDiff, closest)
		}
	} else {
		currDiff := math.Abs(rootValFl - target)
		if currDiff < minDiff {
			return recurse(root.left, target, currDiff, rootValFl)
		} else {
			return recurse(root.left, target, minDiff, closest)
		}
	}
}
