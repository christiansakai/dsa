package solution

func SolveRecursive(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	recurse(root, &result)

	return result
}

func recurse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	recurse(root.Left, result)
	*result = append(*result, root.Val)
	recurse(root.Right, result)
}

func SolveIterative(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	current := root
	stack := []*TreeNode{}

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)
		current = current.Right
	}

	return result
}
