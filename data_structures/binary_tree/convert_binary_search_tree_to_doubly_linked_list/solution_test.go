package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tn1 := &node{val: 1}
	tn2 := &node{val: 2}
	tn3 := &node{val: 3}
	tn4 := &node{val: 4}
	tn5 := &node{val: 5}

	tree := tn4
	tn4.left = tn2
	tn4.right = tn5
	tn2.left = tn1
	tn2.right = tn3

	ln1 := &node{val: 1}
	ln2 := &node{val: 2}
	ln3 := &node{val: 3}
	ln4 := &node{val: 4}
	ln5 := &node{val: 5}

	list := ln1
	ln1.left = ln5
	ln1.right = ln2
	ln2.left = ln1
	ln2.right = ln3
	ln3.left = ln2
	ln3.right = ln4
	ln4.left = ln3
	ln4.right = ln5
	ln5.left = ln4
	ln5.right = ln1

	got := Solve(tree)
	if !reflect.DeepEqual(got, list) {
		t.Errorf("got %v, want %v", got, list)
	}
}
