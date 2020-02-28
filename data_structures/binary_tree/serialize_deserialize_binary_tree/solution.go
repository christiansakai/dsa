package solution

import (
	"strconv"
	"strings"
	// "fmt"
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

	val, _ := strconv.Atoi(strArr[0])
	root := &TreeNode{Val: val}

	recurseDeserialize(strArr, 2, root)
	recurseDeserialize(strArr, 3, root)

	return root
}

func recurseDeserialize(strs []string, index int, root *TreeNode) {
	adjustedIndex := index - 1
	if adjustedIndex >= len(strs) {
		return
	}

	val, _ := strconv.Atoi(strs[adjustedIndex])
	node := &TreeNode{Val: val}

	recurseDeserialize(strs, index*2, node)
	recurseDeserialize(strs, (index*2)+1, node)

	if index%2 == 0 {
		root.Left = node
	} else {
		root.Right = node
	}
}
