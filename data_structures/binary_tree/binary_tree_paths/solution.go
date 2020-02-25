package solution

import (
	"strconv"
	"strings"
)

func Solve(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}

	set := []int{}
	recurse(root, &set, &result)

	return result
}

func recurse(root *TreeNode, set *[]int, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*set = append(*set, root.Val)
		*result = append(*result, toString(set))
		*set = (*set)[:len(*set)-1]

		return
	}

	if root.Left != nil {
		*set = append(*set, root.Val)
		recurse(root.Left, set, result)
		*set = (*set)[:len(*set)-1]
	}

	if root.Right != nil {
		*set = append(*set, root.Val)
		recurse(root.Right, set, result)
		*set = (*set)[:len(*set)-1]
	}
}

func toString(set *[]int) string {
	b := []string{}
	for _, el := range *set {
		b = append(b, strconv.Itoa(el))
	}

	return strings.Join(b, "->")
}
