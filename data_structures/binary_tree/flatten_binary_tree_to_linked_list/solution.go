package solution

func Solve(root *TreeNode) {
	if root == nil {
		return
	}

	recurse(root)
}

func recurse(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := recurse(root.Left)
	right := recurse(root.Right)

	root.Left = nil
	root.Right = left

	tempLeft := left
	for tempLeft.Right != nil {
		tempLeft = tempLeft.Right
	}

	tempLeft.Left = nil
	tempLeft.Right = right

	return root
}
