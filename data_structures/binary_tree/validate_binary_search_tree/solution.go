package solution

func Solve(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	if root.Left != nil && root.Right != nil {
		leftIsValid, _, rightMost := recurse(root.Left)
		rightIsValid, leftMost, _ := recurse(root.Right)

		return leftIsValid && rightIsValid &&
			(root.Val > rightMost) && (root.Val < leftMost)
	}

	if root.Left != nil {
		leftIsValid, _, rightMost := recurse(root.Left)
		return leftIsValid
	}
}

func recurse(root *TreeNode) (bool, int, int) {
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	}

	if root.Left != nil && root.Right != nil {
		currIsValid := (root.Left.Val < root.Val) &&
			(root.Val < root.Right.Val)

		leftIsValid, leftLeftMost, leftRightMost := recurse(root.Left)
		rightIsValid, rightLeftMost, rightRightMost := recurse(root.Right)

		isValid := currIsValid && leftIsValid && rightIsValid &&
			(root.Val > leftRightMost) && (root.Val < rightLeftMost)

		return isValid, leftLeftMost, rightRightMost
	}

	if root.Left != nil {
		currIsValid := root.Left.Val < root.Val
		leftIsValid, leftLeftMost, leftRightMost := recurse(root.Left)

		isValid := currIsValid && leftIsValid && (root.Val > leftRightMost)

		return isValid, leftLeftMost, root.Val
	}

	currIsValid := root.Val > root.Right.Val
	rightIsValid, rightLeftMost, rightRightMost := recurse(root.Right)

	isValid := currIsValid && rightIsValid && (root.Val < rightLeftMost)

	return isValid, root.Val, rightRightMost
}
