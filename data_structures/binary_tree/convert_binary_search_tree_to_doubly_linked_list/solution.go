package solution

func Solve(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	head := recurse(root)
	tail := head

	for tail != nil && tail.Right != nil {
		tail = tail.Right
	}

	head.Left = tail
	tail.Right = head

	return head
}

func recurse(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		subProbLeft := recurse(root.Left)

		tempSubProbLeft := subProbLeft
		for tempSubProbLeft != nil && tempSubProbLeft.Right != nil {
			tempSubProbLeft = tempSubProbLeft.Right
		}

		tempSubProbLeft.Right = root
		root.Left = tempSubProbLeft

		subProbRight := recurse(root.Right)
		root.Right = subProbRight
		subProbRight.Left = root

		return subProbLeft
	} else if root.Left != nil {
		subProbLeft := recurse(root.Left)

		tempSubProbLeft := subProbLeft
		for tempSubProbLeft != nil && tempSubProbLeft.Right != nil {
			tempSubProbLeft = tempSubProbLeft.Right
		}

		tempSubProbLeft.Right = root
		root.Left = tempSubProbLeft

		return subProbLeft
	} else {
		subProbRight := recurse(root.Right)
		root.Right = subProbRight
		subProbRight.Left = root

		return root
	}
}
