package solution

import (
	"strconv"
	"strings"
)

func SolveSerialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []string{}
	recurseSerialize(root, &result)

	return strings.Join(result, ",")
}

func recurseSerialize(root *TreeNode, result *[]string) {
	if root == nil {
		*result = append(*result, "nil")
		return
	}

	*result = append(*result, strconv.Itoa(root.Val))
	recurseSerialize(root.Left, result)
	recurseSerialize(root.Right, result)
}

func SolveDeserialize(str string) *TreeNode {
	if len(str) == 0 {
		return nil
	}

	strArr := strings.Split(str, ",")
	return recurse(&strArr)
}

func recurse(str *[]string) *TreeNode {
	if (*str)[0] == "nil" {
		*str = (*str)[1:]
		return nil
	}

	valStr := (*str)[0]
	*str = (*str)[1:]

	val, _ := strconv.Atoi(valStr)
	root := &TreeNode{
		Val:   val,
		Left:  recurse(str),
		Right: recurse(str),
	}

	return root
}
