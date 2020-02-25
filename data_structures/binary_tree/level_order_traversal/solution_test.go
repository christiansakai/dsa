package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		output [][]int
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			output: [][]int{
				[]int{3},
				[]int{9, 20},
				[]int{15, 7},
			},
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
