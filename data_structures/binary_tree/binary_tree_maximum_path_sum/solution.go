package solution

import "math"

func Solve(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxPathSum, _ := recurse(root)
	return int(maxPathSum)
}

func recurse(root *TreeNode) (float64, float64) {
	if root.Left == nil && root.Right == nil {
		return float64(root.Val), float64(root.Val)
	}

	if root.Left != nil && root.Right != nil {
		leftMaxPathSum, leftMaxBranchSum := recurse(root.Left)
		rightMaxPathSum, rightMaxBranchSum := recurse(root.Right)

		currMaxPathSumWithJustLeft := float64(root.Val) + leftMaxBranchSum
		currMaxPathSumWithJustRight := float64(root.Val) + rightMaxBranchSum
		currMaxPathSum := float64(root.Val) + leftMaxBranchSum + rightMaxBranchSum
		maxPathSum := math.Max(float64(root.Val), currMaxPathSumWithJustLeft)
		maxPathSum = math.Max(float64(root.Val), currMaxPathSumWithJustRight)
		maxPathSum = math.Max(float64(root.Val), currMaxPathSum)

		maxPathSum = math.Max(maxPathSum, leftMaxPathSum)
		maxPathSum = math.Max(maxPathSum, rightMaxPathSum)

		maxBranchSum := math.Max(leftMaxBranchSum, rightMaxBranchSum)

		return maxPathSum, maxBranchSum + float64(root.Val)
	}

	if root.Left != nil {
		leftMaxPathSum, leftMaxBranchSum := recurse(root.Left)

		currMaxPathSum := float64(root.Val) + leftMaxBranchSum

		maxPathSum := math.Max(float64(root.Val), currMaxPathSum)
		maxPathSum = math.Max(maxPathSum, leftMaxPathSum)

		return maxPathSum, leftMaxBranchSum + float64(root.Val)
	}

	rightMaxPathSum, rightMaxBranchSum := recurse(root.Right)

	currMaxPathSum := float64(root.Val) + rightMaxBranchSum

	maxPathSum := math.Max(float64(root.Val), currMaxPathSum)
	maxPathSum = math.Max(maxPathSum, rightMaxPathSum)

	return maxPathSum, rightMaxBranchSum + float64(root.Val)
}
