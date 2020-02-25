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

	*result = append(*result, root.Val)
	recurse(root.Left, result)
	recurse(root.Right, result)
}

func SolveIterative(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
