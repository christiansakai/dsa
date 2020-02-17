package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveSerialize(root *node) string {
	if root == nil {
		return ""
	}

	arr := []string{}
	queue := []*node{root}

	for len(queue) > 0 {
		qLen := len(queue)
		fmt.Println(qLen)
		for i := 0; i < qLen; i++ {
			head := queue[0]
			queue = queue[:len(queue)-1]

			if head == nil {
				arr = append(arr, "nil")
			} else {
				arr = append(arr, strconv.Itoa(head.val))

				queue = append(queue, head.left)
				queue = append(queue, head.right)
			}
		}
	}

	return toString(arr)
}

func toString(arr []string) string {
	str := ""
	for i, el := range arr {
		if i == len(arr)-1 {
			str += el + ","
		} else {
			str += el
		}
	}

	return str
}

func SolveDeserialize(s string) *node {
	if len(s) == 0 {
		return nil
	}

	arr := strings.Split(s, ",")

	return recurse(arr, 0)
}

func recurse(arr []string, index int) *node {
	// fmt.Println(index)
	if index >= len(arr) {
		return nil
	}

	val := arr[index]
	if val == "nil" {
		return nil
	}

	intVal, _ := strconv.Atoi(val)

	return &node{
		val:   intVal,
		left:  recurse(arr, (2*index)+1),
		right: recurse(arr, (2*index)+2),
	}
}
