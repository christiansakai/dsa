package solution

import "testing"

func TestSolution(t *testing.T) {
	tree1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	tree2 := &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	tree3 := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: -1},
	}

	tree4 := &TreeNode{
		Val:  -2,
		Left: &TreeNode{Val: -1},
	}

	tree5 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: -2},
		Right: &TreeNode{Val: 3},
	}

	tests := []struct {
		input  *TreeNode
		output int
	}{
		{tree1, 6},
		{tree2, 42},
		{tree3, 2},
		{tree4, -1},
		{tree5, 4},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}
