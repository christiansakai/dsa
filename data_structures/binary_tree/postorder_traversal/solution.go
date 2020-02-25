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
	recurse(root.Right, result)
	*result = append(*result, root.Val)
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

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(result)

	return result
}

func reverse(result []int) {
	mid := len(result) / 2
	for i := 0; i < mid; i++ {
		temp := result[i]
		result[i] = result[len(result)-1-i]
		result[len(result)-1-i] = temp
	}
}
