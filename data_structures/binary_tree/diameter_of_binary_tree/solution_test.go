package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			output: 3,
		},
	}

	for _, tt := range tests {
		got := Solve(tt.input)

		if got != tt.output {
			t.Errorf("got %d, want %d", got, tt.output)
		}
	}
}
