package solution

import "math"

func Solve(root *TreeNode) int {
  if root == nil {
    return 0
  }

  maxPathSum, _ := recurse(root)
  return maxPathSum
}

func recurse(root *TreeNode) (int, int) {
  if root.Left == nil && root.Right == nil {
    return root.Val, root.Val
  }

  maxPathSum := float64(root.Val)
  longestBranch := math.Inf(-1)

  if root.Left != nil {
    leftMaxPathSum, leftLongestBranch := recurse(root.Left)

    maxPathSum = math.Max(maxPathSum, float64(leftMaxPathSum))
    maxPathSum = math.Max(
      maxPathSum,
      float64(root.Val + leftLongestBranch),
    )

    longestBranch = math.Max(
      longestBranch,
      float64(leftLongestBranch),
    )
  }

  if root.Right != nil {
    rightMaxPathSum, rightLongestBranch := recurse(root.Right)

    maxPathSum = math.Max(maxPathSum, float64(rightMaxPathSum))
    maxPathSum = math.Max(
      maxPathSum,
      float64(root.Val + rightLongestBranch),
    )

    longestBranch = math.Max(
      longestBranch,
      float64(rightLongestBranch),
    )
  }

  return int(maxPathSum), int(longestBranch)
}
