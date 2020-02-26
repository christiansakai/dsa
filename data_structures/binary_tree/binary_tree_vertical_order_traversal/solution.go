package solution

import "sort"

type helper struct {
	level int
	node  *TreeNode
}

func Solve(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Use map here to make easy negative index
	result := map[int][]int{}
	queue := []helper{helper{0, root}}

	for len(queue) > 0 {
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			h := queue[0]
			queue = queue[1:]

			level := h.level
			node := h.node

			if _, ok := result[level]; !ok {
				result[level] = []int{}
			}

			result[level] = append(result[level], node.Val)

			if node.Left != nil {
				queue = append(queue, helper{
					level: level - 1,
					node:  node.Left,
				})
			}

			if node.Right != nil {
				queue = append(queue, helper{
					level: level + 1,
					node:  node.Right,
				})
			}
		}
	}

	return convertToSlice(result)
}

func convertToSlice(result map[int][]int) [][]int {
	keys := []int{}
	for k, _ := range result {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	slice := [][]int{}
	for _, k := range keys {
		slice = append(slice, result[k])
	}

	return slice
}
