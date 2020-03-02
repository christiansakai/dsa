package solution

func Solve(root *TreeNode) {
	if root == nil {
		return
	}

	recurse(root)
}

func recurse(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		subProbLeft := recurse(root.Left)
		subProbRight := recurse(root.Right)

		root.Left = nil
		root.Right = subProbLeft

		// traverse subProbLeft to the end
		for subProbLeft != nil && subProbLeft.Right != nil {
			subProbLeft = subProbLeft.Right
		}

		subProbLeft.Right = subProbRight
	} else if root.Left != nil {
		subProbLeft := recurse(root.Left)
		root.Left = nil
		root.Right = subProbLeft
	} else if root.Right != nil {
		subProbRight := recurse(root.Right)
		root.Left = nil
		root.Right = subProbRight
	}

	return root
}
