package solution

import "testing"

func TestSolution(t *testing.T) {
	t1a := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	t1b := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	t2a := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	t2b := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	t3a := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}

	t3b := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	type input struct {
		ta *TreeNode
		tb *TreeNode
	}
	tests := []struct {
		input  input
		output bool
	}{
		{input{t1a, t1b}, true},
		{input{t2a, t2b}, false},
		{input{t3a, t3b}, false},
	}

	for _, tt := range tests {
		got := Solve(tt.input.ta, tt.input.tb)

		if got != tt.output {
			t.Errorf("got %t, want %t", got, tt.output)
		}
	}
}
