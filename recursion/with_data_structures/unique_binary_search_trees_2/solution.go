package solution

func Solve(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return recurse(1, n)
}

func recurse(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}

	for i := start; i <= end; i++ {
		// All possible left subtrees with i as parent
		leftTrees := recurse(start, i-1)

		// All possible right subtrees with i as parent
		rightTrees := recurse(i+1, end)

		for _, lt := range leftTrees {
			for _, rt := range rightTrees {
				root := &TreeNode{
					Val:   i,
					Left:  lt,
					Right: rt,
				}

				result = append(result, root)
			}
		}
	}

	return result
}
