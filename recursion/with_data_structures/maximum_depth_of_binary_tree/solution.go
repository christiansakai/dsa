package solution

import "math"

func Solve(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := Solve(root.Left)
	right := Solve(root.Right)

	max := int(math.Max(float64(left), float64(right)))

	return 1 + max
}
