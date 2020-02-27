package solution

func Solve(root *TreeNode) {
	if root == nil {
		return
	}

	Solve(root.Left)
	Solve(root.Right)

	tempLeft := root.Left
	tempRight := root.Right

	root.Left = nil
	root.Right = tempLeft

	tempLeft.Left = nil
	tempLeft.Right = tempRight

	tempRight.Left = nil
	tempRight.Right = nil
}
