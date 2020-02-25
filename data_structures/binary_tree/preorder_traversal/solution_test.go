package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output []int
	}{
		{
			input: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			output: []int{1, 2, 3},
		},
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			output: []int{3, 1, 2},
		},
	}

	t.Run("Recursive", func(t *testing.T) {
		for _, tt := range tests {
			got := SolveRecursive(tt.input)

			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		}
	})

	t.Run("Iterative", func(t *testing.T) {
		for _, tt := range tests {
			got := SolveIterative(tt.input)

			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		}
	})
}
