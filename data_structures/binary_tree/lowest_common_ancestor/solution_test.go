package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	t0 := &TreeNode{Val: 0}
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 6}
	t7 := &TreeNode{Val: 7}
	t8 := &TreeNode{Val: 8}

	t3.Left = t5
	t3.Right = t1

	t5.Left = t6
	t5.Right = t2

	t2.Left = t7
	t2.Right = t4

	t1.Left = t0
	t1.Right = t8

	tests := []struct {
		input struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}
		output *TreeNode
	}{
		{struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{t3, t5, t1}, t3},
		{struct {
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
		}{t3, t5, t4}, t5},
	}

	for _, tt := range tests {
		got := Solve(tt.input.root, tt.input.p, tt.input.q)

		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("got %v, want %v", got, tt.output)
		}
	}
}
