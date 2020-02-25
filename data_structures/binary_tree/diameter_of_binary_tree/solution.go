package solution

import "math"

func Solve(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter, _ := recurse(root)
	return int(maxDiameter)
}

func recurse(root *TreeNode) (float64, float64) {
	if root.Left == nil && root.Right == nil {
		return 0, 0
	}

	if root.Left != nil && root.Right != nil {
		leftMaxDiameter, leftMaxBranchLength := recurse(root.Left)
		rightMaxDiameter, rightMaxBranchLength := recurse(root.Right)

		currDiameter := 2 + leftMaxBranchLength + rightMaxBranchLength

		maxDiameter := currDiameter
		maxDiameter = math.Max(maxDiameter, leftMaxDiameter)
		maxDiameter = math.Max(maxDiameter, rightMaxDiameter)

		maxBranchLength := math.Max(
			leftMaxBranchLength,
			rightMaxBranchLength,
		)

		return maxDiameter, maxBranchLength + 1
	}

	if root.Left != nil {
		leftMaxDiameter, leftMaxBranchLength := recurse(root.Left)

		currDiameter := 1 + leftMaxBranchLength

		maxDiameter := math.Max(currDiameter, leftMaxDiameter)

		return maxDiameter, leftMaxBranchLength + 1
	}

	rightMaxDiameter, rightMaxBranchLength := recurse(root.Right)

	currDiameter := 1 + rightMaxBranchLength

	maxDiameter := math.Max(currDiameter, rightMaxDiameter)

	return maxDiameter, rightMaxBranchLength + 1
}
