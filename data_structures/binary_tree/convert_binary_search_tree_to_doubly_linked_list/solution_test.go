package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tn1 := &TreeNode{Val: 1}
	tn2 := &TreeNode{Val: 2}
	tn3 := &TreeNode{Val: 3}
	tn4 := &TreeNode{Val: 4}
	tn5 := &TreeNode{Val: 5}

	tree := tn4
	tn4.Left = tn2
	tn4.Right = tn5
	tn2.Left = tn1
	tn2.Right = tn3

	ln1 := &TreeNode{Val: 1}
	ln2 := &TreeNode{Val: 2}
	ln3 := &TreeNode{Val: 3}
	ln4 := &TreeNode{Val: 4}
	ln5 := &TreeNode{Val: 5}

	list := ln1
	ln1.Left = ln5
	ln1.Right = ln2
	ln2.Left = ln1
	ln2.Right = ln3
	ln3.Left = ln2
	ln3.Right = ln4
	ln4.Left = ln3
	ln4.Right = ln5
	ln5.Left = ln4
	ln5.Right = ln1

	got := Solve(tree)
	if !reflect.DeepEqual(got, list) {
		t.Errorf("got %v, want %v", got, list)
	}
}
