package solution

func Solve(root *node) *node {
	if root == nil {
		return nil
	}

	arr := []int{}
	collectValInorder(root, &arr)

	return transformToDLL(arr)
}

func collectValInorder(n *node, arr *[]int) {
	if n == nil {
		return
	}

	collectValInorder(n.left, arr)
	*arr = append(*arr, n.val)
	collectValInorder(n.right, arr)
}

func transformToDLL(arr []int) *node {
	head := &node{val: arr[0]}
	prev := head

	for i := 1; i < len(arr); i++ {
		newNode := &node{val: arr[i]}

		newNode.left = prev
		prev.right = newNode

		prev = newNode
	}

	prev.right = head
	head.left = prev

	return head
}
