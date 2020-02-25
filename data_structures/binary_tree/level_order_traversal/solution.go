package solution

func SolveRecursive(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	level := 0
	recurse(root, level, &result)

	return result
}

func recurse(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result)-1 < level {
		newLevel := []int{}
		*result = append(*result, newLevel)
	}

	(*result)[level] = append((*result)[level], root.Val)
	recurse(root.Left, level+1, result)
	recurse(root.Right, level+1, result)
}

func SolveIterative(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []int{}
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}
